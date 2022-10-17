package compiler

import (
	"flag"
	"fmt"
	"functional-language-compiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Compiler struct {
	visitor                *BuildASTVisitor
	typeChecker            *TypeChecker
	codeGenerator          *CodeGenerator
	compiledFiles          map[string]void
	filesDuringCompilation map[string]void
	codeDir                string
	executionDir           string
	cmd                    string
	inputFile              string
	outputFile             string
	target                 string
	tags                   []string
}

func NewCompiler() *Compiler {
	c := new(Compiler)
	c.init()
	return c
}

func (c *Compiler) init() {
	c.visitor = new(BuildASTVisitor)

	c.typeChecker = new(TypeChecker)
	c.typeChecker.init()

	c.codeGenerator = newCodeGenerator()

	c.compiledFiles = make(map[string]void)
	c.filesDuringCompilation = make(map[string]void)

	_, codePath, _, _ := runtime.Caller(0)
	c.codeDir = filepath.Dir(filepath.Dir(codePath))

	executionPath, _ := os.Getwd()
	c.executionDir = executionPath
}

func (c *Compiler) Compile() {
	gCodeDirectory := filepath.Join(c.codeDir, "gmachinerun")
	gCodeFile := filepath.Join(gCodeDirectory, "main.go")
	c.readCommand()
	c.addCompiledGlobalsTypes("compilerInternals.hs")

	switch c.target {
	case linuxTargetName:
		c.addCompiledGlobalsTypes("compiledGlobalsLinux.hs")
	case featherTargetName:
		c.addCompiledGlobalsTypes("compiledGlobalsFeather.hs")
	default:
		c.addCompiledGlobalsTypes("compiledGlobals.hs")
	}

	c.compileAllToGCode(c.inputFile)
	c.emitGCode(gCodeFile)
	c.compileGCode(gCodeDirectory)
}

func (c *Compiler) readCommand() {
	buildCmd := flag.NewFlagSet(buildCmdName, flag.ExitOnError)
	flashCmd := flag.NewFlagSet(flashCmdName, flag.ExitOnError)

	buildTarget := buildCmd.String("target", linuxTargetName, "target")
	buildOutput := buildCmd.String("o", "a", "output file")

	flashTarget := flashCmd.String("target", "", "target")

	if len(os.Args) < 2 {
		fmt.Printf("expected '%s' or '%s' subcommands", buildCmdName, flashCmdName)
		os.Exit(1)
	}

	switch os.Args[1] {
	case buildCmdName:
		buildCmd.Parse(os.Args[2:])

		c.cmd = buildCmdName
		c.target = *buildTarget
		c.inputFile = buildCmd.Arg(0)

		if filepath.IsAbs(*buildOutput) {
			c.outputFile = *buildOutput
		} else {
			c.outputFile = filepath.Join(c.executionDir, *buildOutput)
		}
	case flashCmdName:
		flashCmd.Parse(os.Args[2:])
		//asdf
		c.cmd = flashCmdName
		c.target = *flashTarget
		c.inputFile = flashCmd.Arg(0)
	}

	switch c.target {
	case linuxTargetName:
		c.tags = append(c.tags, linuxTargetTag)
	case featherTargetName:
		c.tags = append(c.tags, featherTargetTag)
	}
}

const (
	buildCmdName = "build"
	flashCmdName = "flash"
)

const (
	linuxTargetName   = "linux"
	featherTargetName = "feather"
)

const (
	linuxTargetTag   = "lin"
	featherTargetTag = "feather"
)

func (c *Compiler) addCompiledGlobalsTypes(filename string) {
	c.compileSubtreeToGCode(filename)
}

func (c *Compiler) compileAllToGCode(filename string) {
	if filename == "" {
		errFatalAnywhere("Input file not provided")
	}

	c.compileSubtreeToGCode(filename)

	c.typeChecker.generalCheck()
}

func (c *Compiler) compileSubtreeToGCode(filename string) {
	_, duringCompilation := c.filesDuringCompilation[filename]
	if duringCompilation {
		errFatalAnywhere(fmt.Sprintf("Circular dependency in `%s`", filename))
	}

	_, compiled := c.compiledFiles[filename]
	if compiled {
		return
	}

	c.filesDuringCompilation[filename] = voidMember

	var absPath string
	if filepath.IsAbs(filename) {
		absPath = filename
	} else {
		absPath = filepath.Join(c.codeDir, "haskellPredefined", filename)
		if _, err := os.Stat(absPath); err != nil {
			absPath = filepath.Join(c.executionDir, filename)
		}
	}

	ast := c.parseFile(absPath)
	c.compileImportedFiles(ast)
	c.typeChecker.checkCodePart(ast)
	c.codeGenerator.genCodePart(ast)

	delete(c.filesDuringCompilation, filename)
	c.compiledFiles[filename] = voidMember
}

func (c *Compiler) emitGCode(outputFile string) {
	c.codeGenerator.emit(outputFile)
}

func (c *Compiler) compileGCode(inputDirectory string) {
	var args []string
	tags := strings.Join(c.tags, " ")
	relInputDir, _ := filepath.Rel(c.executionDir, inputDirectory)
	relInputDir = "./" + relInputDir

	//args = append(args, "tinygo")
	args = append(args, c.cmd)
	args = append(args, fmt.Sprintf("-target=%s", c.getTinyGoTarget()))
	args = append(args, fmt.Sprintf("-tags=%q", tags))
	//args = append(args, "-size", "full")

	if c.cmd == buildCmdName {
		args = append(args, "-o", c.outputFile)
	}

	args = append(args, relInputDir)

	exec.Command("go", "env", "-w", "GO111MODULE=auto").Run()
	cmd := exec.Command("tinygo", args...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(stdout))
}

func (c *Compiler) parseFile(filename string) *TopDeclsList { // Takes absolute path to the file
	errorListener := &ParserErrorListener{}

	// Setup the input
	is, err := antlr.NewFileStream(filename)
	if err != nil {
		errFatalAnywhere(fmt.Sprintf("File '%s' not found", filename))
	}

	// Create the Lexer
	lexer := parser.NewLanguageLexer(is)
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewLanguageParser(stream)
	p.AddErrorListener(errorListener)

	ast := p.Start().Accept(c.visitor).(*TopDeclsList)

	if errorListener.errorsCount > 0 {
		os.Exit(1)
	}

	return ast
}

func (c *Compiler) compileImportedFiles(ast *TopDeclsList) {
	for _, decl := range ast.decls {
		importDecl, ok := decl.(*ImportTopDecl)
		if ok {
			importedFile := importDecl.file
			c.compileSubtreeToGCode(importedFile)
		}
	}
}

func (c *Compiler) getTinyGoTarget() string {
	switch c.target {
	case featherTargetName:
		return "feather-nrf52840-sense"
	default:
		return ""
	}
}

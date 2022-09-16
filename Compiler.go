package main

import (
	"fmt"
	"functional-language-compiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"os/exec"
)

type Compiler struct {
	visitor                *BuildASTVisitor
	typeChecker            *TypeChecker
	codeGenerator          *CodeGenerator
	compiledFiles          map[string]void
	filesDuringCompilation map[string]void
}

func newCompiler() *Compiler {
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
}

func (c *Compiler) compile(filename string) {
	gCodeDirectory := "./gmachinerun/"
	gCodeFile := gCodeDirectory + "main.go"
	c.addCompiledGlobalsTypes("./compiledGlobals.hs")
	c.compileAllToGCode(filename)
	c.emitGCode(gCodeFile)
	c.compileGCode(gCodeDirectory)
}

func (c *Compiler) addCompiledGlobalsTypes(filename string) {
	ast := c.parseFile(filename)
	c.typeChecker.checkCodePart(ast)
}

func (c *Compiler) compileAllToGCode(filename string) {
	c.compileSubtreeToGCode(filename)

	// TODO: check things that should apply to the whole code (like main function existence)
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

	ast := c.parseFile(filename)
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
	cmd := exec.Command("tinygo", "build", "-size", "full", "-o", "./gmachinerun/output", inputDirectory)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(stdout))
}

func (c *Compiler) parseFile(filename string) *TopDeclsList {
	errorListener := &ParserErrorListener{}

	// Setup the input
	is, _ := antlr.NewFileStream(filename)

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

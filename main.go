package main

import (
	"fmt"
	"functional-language-compiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/davecgh/go-spew/spew"
	"os"
	"os/exec"
)

func main() {
	// spew config
	//spew.Config.DisablePointerAddresses = true
	spew.Config.Indent = "  "

	// Setup the input
	is, _ := antlr.NewFileStream("input.hs")
	errorListener := &ParserErrorListener{}
	//is := antlr.NewInputStream("data Maybe a = Just a | Nothing; int5 :: Maybe Int; int5 = Nothing")

	// Create the Lexer
	lexer := parser.NewLanguageLexer(is)
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewLanguageParser(stream)
	p.AddErrorListener(errorListener)

	// Finally parse the expression
	visitor := new(BuildASTVisitor)

	ast := p.Start().Accept(visitor).(*TopDeclsList)

	if errorListener.errorsCount > 0 {
		os.Exit(1)
	}

	typeChecker := new(TypeChecker)
	typeChecker.init()
	typeChecker.checkTopDecls(ast)

	codeGenerator := newCodeGenerator()
	codeGenerator.gen(ast)

	cmd := exec.Command("tinygo", "build", "-size", "full", "-o", "./gmachinerun/output", "./gmachinerun/")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(stdout))
	//spew.Dump(typeChecker)
	//spew.Dump(ast)
}

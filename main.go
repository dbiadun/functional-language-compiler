package main

import (
	"functional-language-compiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// spew config
	spew.Config.DisablePointerAddresses = true
	spew.Config.Indent = "  "

	// Setup the input
	is := antlr.NewInputStream("five = 5")

	// Create the Lexer
	lexer := parser.NewLanguageLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewLanguageParser(stream)

	// Finally parse the expression
	visitor := new(BuildASTVisitor)

	ast := p.Start().Accept(visitor).(*TopDeclsList)

	spew.Dump(ast)
}

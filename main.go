package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/davecgh/go-spew/spew"
	"participle-test/parser"
)

type languageListener struct {
	*parser.BaseLanguageListener
}

func main() {
	// spew config
	spew.Config.DisablePointerAddresses = true
	spew.Config.Indent = "  "

	// Setup the input
	is := antlr.NewInputStream("1 / 2 * 3")

	// Create the Lexer
	lexer := parser.NewLanguageLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewLanguageParser(stream)

	// Finally parse the expression
	visitor := BuildASTVisitor{}

	x := p.Start().Accept(&visitor)
	spew.Dump(x)
}

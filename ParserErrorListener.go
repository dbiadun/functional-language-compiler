package main

import "github.com/antlr/antlr4/runtime/Go/antlr"

type ParserErrorListener struct {
	*antlr.DefaultErrorListener
	errorsCount int
}

func (l *ParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errorsCount++
}

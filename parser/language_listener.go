// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LanguageListener is a complete listener for a parse tree produced by LanguageParser.
type LanguageListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterEMulDiv is called when entering the EMulDiv production.
	EnterEMulDiv(c *EMulDivContext)

	// EnterEAddSub is called when entering the EAddSub production.
	EnterEAddSub(c *EAddSubContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitEMulDiv is called when exiting the EMulDiv production.
	ExitEMulDiv(c *EMulDivContext)

	// ExitEAddSub is called when exiting the EAddSub production.
	ExitEAddSub(c *EAddSubContext)
}

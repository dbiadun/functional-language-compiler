// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLanguageListener is a complete listener for a parse tree produced by LanguageParser.
type BaseLanguageListener struct{}

var _ LanguageListener = &BaseLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseLanguageListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseLanguageListener) ExitStart(ctx *StartContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseLanguageListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseLanguageListener) ExitNumber(ctx *NumberContext) {}

// EnterEMulDiv is called when production EMulDiv is entered.
func (s *BaseLanguageListener) EnterEMulDiv(ctx *EMulDivContext) {}

// ExitEMulDiv is called when production EMulDiv is exited.
func (s *BaseLanguageListener) ExitEMulDiv(ctx *EMulDivContext) {}

// EnterEAddSub is called when production EAddSub is entered.
func (s *BaseLanguageListener) EnterEAddSub(ctx *EAddSubContext) {}

// ExitEAddSub is called when production EAddSub is exited.
func (s *BaseLanguageListener) ExitEAddSub(ctx *EAddSubContext) {}

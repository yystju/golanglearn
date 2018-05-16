// Code generated from Test.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Test

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTestListener is a complete listener for a parse tree produced by TestParser.
type BaseTestListener struct{}

var _ TestListener = &BaseTestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTestListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTestListener) ExitStart(ctx *StartContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseTestListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseTestListener) ExitExp(ctx *ExpContext) {}

// EnterSub is called when production sub is entered.
func (s *BaseTestListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BaseTestListener) ExitSub(ctx *SubContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseTestListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseTestListener) ExitExpr(ctx *ExprContext) {}

// EnterExp0 is called when production exp0 is entered.
func (s *BaseTestListener) EnterExp0(ctx *Exp0Context) {}

// ExitExp0 is called when production exp0 is exited.
func (s *BaseTestListener) ExitExp0(ctx *Exp0Context) {}

// EnterExp1 is called when production exp1 is entered.
func (s *BaseTestListener) EnterExp1(ctx *Exp1Context) {}

// ExitExp1 is called when production exp1 is exited.
func (s *BaseTestListener) ExitExp1(ctx *Exp1Context) {}

// EnterExp2 is called when production exp2 is entered.
func (s *BaseTestListener) EnterExp2(ctx *Exp2Context) {}

// ExitExp2 is called when production exp2 is exited.
func (s *BaseTestListener) ExitExp2(ctx *Exp2Context) {}

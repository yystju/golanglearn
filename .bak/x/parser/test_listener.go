// Code generated from Test.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Test

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TestListener is a complete listener for a parse tree produced by TestParser.
type TestListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterSub is called when entering the sub production.
	EnterSub(c *SubContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExp0 is called when entering the exp0 production.
	EnterExp0(c *Exp0Context)

	// EnterExp1 is called when entering the exp1 production.
	EnterExp1(c *Exp1Context)

	// EnterExp2 is called when entering the exp2 production.
	EnterExp2(c *Exp2Context)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitSub is called when exiting the sub production.
	ExitSub(c *SubContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExp0 is called when exiting the exp0 production.
	ExitExp0(c *Exp0Context)

	// ExitExp1 is called when exiting the exp1 production.
	ExitExp1(c *Exp1Context)

	// ExitExp2 is called when exiting the exp2 production.
	ExitExp2(c *Exp2Context)
}

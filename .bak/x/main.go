package main

import (
	"./parser"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TestListenerType struct {
	*parser.BaseTestListener
	stack []interface{}
}

func (t *TestListenerType) push(i interface{}) {
	fmt.Println("[PUSH] ", i)
	t.stack = append(t.stack, i)
}

func (t *TestListenerType) pop() interface{} {
	fmt.Println("[POP]")

	if len(t.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	result := t.stack[len(t.stack)-1]

	t.stack = t.stack[:len(t.stack)-1]

	fmt.Println("[POP] ", result)

	return result
}

func (t *TestListenerType) EnterExp(ctx *parser.ExpContext) {
	fmt.Println("[EnterExp]")
	switch ctx.GetCmd().GetTokenType() {
	case parser.TestParserPLUS:
		fmt.Println("+")
	case parser.TestParserMINUS:
		fmt.Println("-")
	}
}

func (t *TestListenerType) ExitExp(ctx *parser.ExpContext) {

}

func (t *TestListenerType) EnterSub(ctx *parser.SubContext) {
	fmt.Println("[EnterSub]", ctx.GetText())
	t.push(ctx.GetText())
}

func (t *TestListenerType) ExitSub(ctx *parser.SubContext) {
	fmt.Println("[ExitSub]", ctx.GetText())

	o := t.pop().(string)

	fmt.Println(o)
}

func (t *TestListenerType) ExitExp1(ctx *parser.Exp1Context) {
	fmt.Println("1", ctx.GetText())
}

func (t *TestListenerType) ExitExp2(ctx *parser.Exp2Context) {
	fmt.Println("2", ctx.GetText())
}

func eval(s string) {
	i := antlr.NewInputStream(s)
	l := parser.NewTestLexer(i)
	r := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)
	p := parser.NewTestParser(r)

	var n *TestListenerType = new(TestListenerType)

	antlr.ParseTreeWalkerDefault.Walk(n, p.Start())
}

var (
	exp string
)

func init() {
	flag.StringVar(&exp, "e", "A;C;D;", "The expression to evaluate.")
	flag.Parse()
}

func main() {
	eval(exp)
}

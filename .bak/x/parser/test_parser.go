// Code generated from Test.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Test

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 52, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 6, 2, 18, 10, 2, 13, 2, 14, 2, 19, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 5, 3, 27, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7,
	5, 36, 10, 5, 12, 5, 14, 5, 39, 11, 5, 3, 5, 5, 5, 42, 10, 5, 3, 6, 3,
	6, 5, 6, 46, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8,
	10, 12, 14, 2, 5, 3, 2, 13, 14, 3, 2, 4, 6, 3, 2, 7, 9, 2, 49, 2, 17, 3,
	2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 28, 3, 2, 2, 2, 8, 32, 3, 2, 2, 2, 10, 45,
	3, 2, 2, 2, 12, 47, 3, 2, 2, 2, 14, 49, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2,
	17, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3,
	2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 7, 2, 2, 3, 22, 3, 3, 2, 2, 2, 23,
	24, 9, 2, 2, 2, 24, 27, 5, 8, 5, 2, 25, 27, 5, 6, 4, 2, 26, 23, 3, 2, 2,
	2, 26, 25, 3, 2, 2, 2, 27, 5, 3, 2, 2, 2, 28, 29, 7, 11, 2, 2, 29, 30,
	5, 8, 5, 2, 30, 31, 7, 12, 2, 2, 31, 7, 3, 2, 2, 2, 32, 37, 5, 10, 6, 2,
	33, 34, 7, 3, 2, 2, 34, 36, 5, 10, 6, 2, 35, 33, 3, 2, 2, 2, 36, 39, 3,
	2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39,
	37, 3, 2, 2, 2, 40, 42, 7, 3, 2, 2, 41, 40, 3, 2, 2, 2, 41, 42, 3, 2, 2,
	2, 42, 9, 3, 2, 2, 2, 43, 46, 5, 12, 7, 2, 44, 46, 5, 14, 8, 2, 45, 43,
	3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 11, 3, 2, 2, 2, 47, 48, 9, 3, 2, 2,
	48, 13, 3, 2, 2, 2, 49, 50, 9, 4, 2, 2, 50, 15, 3, 2, 2, 2, 7, 19, 26,
	37, 41, 45,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'A'", "'B'", "'C'", "'D'", "'E'", "'F'", "", "'('", "')'",
	"'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "WHITESPACE", "LB", "RB", "PLUS", "MINUS",
}

var ruleNames = []string{
	"start", "exp", "sub", "expr", "exp0", "exp1", "exp2",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TestParser struct {
	*antlr.BaseParser
}

func NewTestParser(input antlr.TokenStream) *TestParser {
	this := new(TestParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Test.g4"

	return this
}

// TestParser tokens.
const (
	TestParserEOF        = antlr.TokenEOF
	TestParserT__0       = 1
	TestParserT__1       = 2
	TestParserT__2       = 3
	TestParserT__3       = 4
	TestParserT__4       = 5
	TestParserT__5       = 6
	TestParserT__6       = 7
	TestParserWHITESPACE = 8
	TestParserLB         = 9
	TestParserRB         = 10
	TestParserPLUS       = 11
	TestParserMINUS      = 12
)

// TestParser rules.
const (
	TestParserRULE_start = 0
	TestParserRULE_exp   = 1
	TestParserRULE_sub   = 2
	TestParserRULE_expr  = 3
	TestParserRULE_exp0  = 4
	TestParserRULE_exp1  = 5
	TestParserRULE_exp2  = 6
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TestParserEOF, 0)
}

func (s *StartContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *StartContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TestParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TestParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TestParserLB)|(1<<TestParserPLUS)|(1<<TestParserMINUS))) != 0) {
		{
			p.SetState(14)
			p.Exp()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
		p.Match(TestParserEOF)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCmd returns the cmd token.
	GetCmd() antlr.Token

	// SetCmd sets the cmd token.
	SetCmd(antlr.Token)

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cmd    antlr.Token
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetCmd() antlr.Token { return s.cmd }

func (s *ExpContext) SetCmd(v antlr.Token) { s.cmd = v }

func (s *ExpContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TestParserPLUS, 0)
}

func (s *ExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TestParserMINUS, 0)
}

func (s *ExpContext) Sub() ISubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *TestParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TestParserRULE_exp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TestParserPLUS, TestParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpContext).cmd = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TestParserPLUS || _la == TestParserMINUS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpContext).cmd = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(22)
			p.Expr()
		}

	case TestParserLB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Sub()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubContext is an interface to support dynamic dispatch.
type ISubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubContext differentiates from other interfaces.
	IsSubContext()
}

type SubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubContext() *SubContext {
	var p = new(SubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_sub
	return p
}

func (*SubContext) IsSubContext() {}

func NewSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubContext {
	var p = new(SubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_sub

	return p
}

func (s *SubContext) GetParser() antlr.Parser { return s.parser }

func (s *SubContext) LB() antlr.TerminalNode {
	return s.GetToken(TestParserLB, 0)
}

func (s *SubContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubContext) RB() antlr.TerminalNode {
	return s.GetToken(TestParserRB, 0)
}

func (s *SubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterSub(s)
	}
}

func (s *SubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitSub(s)
	}
}

func (p *TestParser) Sub() (localctx ISubContext) {
	localctx = NewSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TestParserRULE_sub)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(TestParserLB)
	}
	{
		p.SetState(27)
		p.Expr()
	}
	{
		p.SetState(28)
		p.Match(TestParserRB)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExp0() []IExp0Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExp0Context)(nil)).Elem())
	var tst = make([]IExp0Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExp0Context)
		}
	}

	return tst
}

func (s *ExprContext) Exp0(i int) IExp0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp0Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExp0Context)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *TestParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TestParserRULE_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Exp0()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(31)
				p.Match(TestParserT__0)
			}
			{
				p.SetState(32)
				p.Exp0()
			}

		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TestParserT__0 {
		{
			p.SetState(38)
			p.Match(TestParserT__0)
		}

	}

	return localctx
}

// IExp0Context is an interface to support dynamic dispatch.
type IExp0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp0Context differentiates from other interfaces.
	IsExp0Context()
}

type Exp0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp0Context() *Exp0Context {
	var p = new(Exp0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_exp0
	return p
}

func (*Exp0Context) IsExp0Context() {}

func NewExp0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp0Context {
	var p = new(Exp0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_exp0

	return p
}

func (s *Exp0Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp0Context) Exp1() IExp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp1Context)
}

func (s *Exp0Context) Exp2() IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
}

func (s *Exp0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterExp0(s)
	}
}

func (s *Exp0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitExp0(s)
	}
}

func (p *TestParser) Exp0() (localctx IExp0Context) {
	localctx = NewExp0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TestParserRULE_exp0)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TestParserT__1, TestParserT__2, TestParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Exp1()
		}

	case TestParserT__4, TestParserT__5, TestParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Exp2()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExp1Context is an interface to support dynamic dispatch.
type IExp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp1Context differentiates from other interfaces.
	IsExp1Context()
}

type Exp1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp1Context() *Exp1Context {
	var p = new(Exp1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_exp1
	return p
}

func (*Exp1Context) IsExp1Context() {}

func NewExp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp1Context {
	var p = new(Exp1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_exp1

	return p
}

func (s *Exp1Context) GetParser() antlr.Parser { return s.parser }
func (s *Exp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterExp1(s)
	}
}

func (s *Exp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitExp1(s)
	}
}

func (p *TestParser) Exp1() (localctx IExp1Context) {
	localctx = NewExp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TestParserRULE_exp1)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TestParserT__1)|(1<<TestParserT__2)|(1<<TestParserT__3))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExp2Context is an interface to support dynamic dispatch.
type IExp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp2Context differentiates from other interfaces.
	IsExp2Context()
}

type Exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp2Context() *Exp2Context {
	var p = new(Exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TestParserRULE_exp2
	return p
}

func (*Exp2Context) IsExp2Context() {}

func NewExp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2Context {
	var p = new(Exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TestParserRULE_exp2

	return p
}

func (s *Exp2Context) GetParser() antlr.Parser { return s.parser }
func (s *Exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.EnterExp2(s)
	}
}

func (s *Exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TestListener); ok {
		listenerT.ExitExp2(s)
	}
}

func (p *TestParser) Exp2() (localctx IExp2Context) {
	localctx = NewExp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TestParserRULE_exp2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TestParserT__4)|(1<<TestParserT__5)|(1<<TestParserT__6))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

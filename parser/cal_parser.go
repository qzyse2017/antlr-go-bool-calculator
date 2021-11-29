// Code generated from Cal.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Cal

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 33, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 15, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 20, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 28, 10, 4, 12, 4, 14, 4, 31, 11, 4, 3, 4, 2, 3, 6, 5,
	2, 4, 6, 2, 2, 2, 33, 2, 8, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 19, 3, 2,
	2, 2, 8, 9, 5, 6, 4, 2, 9, 3, 3, 2, 2, 2, 10, 15, 7, 3, 2, 2, 11, 12, 7,
	3, 2, 2, 12, 13, 7, 7, 2, 2, 13, 15, 7, 3, 2, 2, 14, 10, 3, 2, 2, 2, 14,
	11, 3, 2, 2, 2, 15, 5, 3, 2, 2, 2, 16, 17, 8, 4, 1, 2, 17, 20, 7, 4, 2,
	2, 18, 20, 5, 4, 3, 2, 19, 16, 3, 2, 2, 2, 19, 18, 3, 2, 2, 2, 20, 29,
	3, 2, 2, 2, 21, 22, 12, 4, 2, 2, 22, 23, 7, 6, 2, 2, 23, 28, 5, 6, 4, 5,
	24, 25, 12, 3, 2, 2, 25, 26, 7, 5, 2, 2, 26, 28, 5, 6, 4, 4, 27, 21, 3,
	2, 2, 2, 27, 24, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29,
	30, 3, 2, 2, 2, 30, 7, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 6, 14, 19, 27, 29,
}
var literalNames []string

var symbolicNames = []string{
	"", "MYINT", "NAME", "OR", "AND", "NUMCOMP", "STRINGCOMP", "WS",
}

var ruleNames = []string{
	"userexpr", "intexpr", "expr",
}

type CalParser struct {
	*antlr.BaseParser
}

// NewCalParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CalParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalParser(input antlr.TokenStream) *CalParser {
	this := new(CalParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Cal.g4"

	return this
}

// CalParser tokens.
const (
	CalParserEOF        = antlr.TokenEOF
	CalParserMYINT      = 1
	CalParserNAME       = 2
	CalParserOR         = 3
	CalParserAND        = 4
	CalParserNUMCOMP    = 5
	CalParserSTRINGCOMP = 6
	CalParserWS         = 7
)

// CalParser rules.
const (
	CalParserRULE_userexpr = 0
	CalParserRULE_intexpr  = 1
	CalParserRULE_expr     = 2
)

// IUserexprContext is an interface to support dynamic dispatch.
type IUserexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUserexprContext differentiates from other interfaces.
	IsUserexprContext()
}

type UserexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserexprContext() *UserexprContext {
	var p = new(UserexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalParserRULE_userexpr
	return p
}

func (*UserexprContext) IsUserexprContext() {}

func NewUserexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserexprContext {
	var p = new(UserexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalParserRULE_userexpr

	return p
}

func (s *UserexprContext) GetParser() antlr.Parser { return s.parser }

func (s *UserexprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UserexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalListener); ok {
		listenerT.EnterUserexpr(s)
	}
}

func (s *UserexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalListener); ok {
		listenerT.ExitUserexpr(s)
	}
}

func (p *CalParser) Userexpr() (localctx IUserexprContext) {
	localctx = NewUserexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalParserRULE_userexpr)

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
		p.SetState(6)
		p.expr(0)
	}

	return localctx
}

// IIntexprContext is an interface to support dynamic dispatch.
type IIntexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntexprContext differentiates from other interfaces.
	IsIntexprContext()
}

type IntexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntexprContext() *IntexprContext {
	var p = new(IntexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalParserRULE_intexpr
	return p
}

func (*IntexprContext) IsIntexprContext() {}

func NewIntexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntexprContext {
	var p = new(IntexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalParserRULE_intexpr

	return p
}

func (s *IntexprContext) GetParser() antlr.Parser { return s.parser }

func (s *IntexprContext) AllMYINT() []antlr.TerminalNode {
	return s.GetTokens(CalParserMYINT)
}

func (s *IntexprContext) MYINT(i int) antlr.TerminalNode {
	return s.GetToken(CalParserMYINT, i)
}

func (s *IntexprContext) NUMCOMP() antlr.TerminalNode {
	return s.GetToken(CalParserNUMCOMP, 0)
}

func (s *IntexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalListener); ok {
		listenerT.EnterIntexpr(s)
	}
}

func (s *IntexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalListener); ok {
		listenerT.ExitIntexpr(s)
	}
}

func (p *CalParser) Intexpr() (localctx IIntexprContext) {
	localctx = NewIntexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalParserRULE_intexpr)

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

	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.Match(CalParserMYINT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Match(CalParserMYINT)
		}
		{
			p.SetState(10)
			p.Match(CalParserNUMCOMP)
		}
		{
			p.SetState(11)
			p.Match(CalParserMYINT)
		}

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
	p.RuleIndex = CalParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) NAME() antlr.TerminalNode {
	return s.GetToken(CalParserNAME, 0)
}

func (s *ExprContext) Intexpr() IIntexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntexprContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(CalParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(CalParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *CalParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CalParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CalParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalParserNAME:
		{
			p.SetState(15)
			p.Match(CalParserNAME)
		}

	case CalParserMYINT:
		{
			p.SetState(16)
			p.Intexpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalParserRULE_expr)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(20)
					p.Match(CalParserAND)
				}
				{
					p.SetState(21)
					p.expr(3)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalParserRULE_expr)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(23)
					p.Match(CalParserOR)
				}
				{
					p.SetState(24)
					p.expr(2)
				}

			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

func (p *CalParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

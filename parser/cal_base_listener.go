// Code generated from Cal.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Cal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalListener is a complete listener for a parse tree produced by CalParser.
type BaseCalListener struct{}

var _ CalListener = &BaseCalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUserexpr is called when production userexpr is entered.
func (s *BaseCalListener) EnterUserexpr(ctx *UserexprContext) {}

// ExitUserexpr is called when production userexpr is exited.
func (s *BaseCalListener) ExitUserexpr(ctx *UserexprContext) {}

// EnterIntexpr is called when production intexpr is entered.
func (s *BaseCalListener) EnterIntexpr(ctx *IntexprContext) {}

// ExitIntexpr is called when production intexpr is exited.
func (s *BaseCalListener) ExitIntexpr(ctx *IntexprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseCalListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseCalListener) ExitExpr(ctx *ExprContext) {}

// Code generated from Cal.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Cal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalListener is a complete listener for a parse tree produced by CalParser.
type CalListener interface {
	antlr.ParseTreeListener

	// EnterUserexpr is called when entering the userexpr production.
	EnterUserexpr(c *UserexprContext)

	// EnterIntexpr is called when entering the intexpr production.
	EnterIntexpr(c *IntexprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitUserexpr is called when exiting the userexpr production.
	ExitUserexpr(c *UserexprContext)

	// ExitIntexpr is called when exiting the intexpr production.
	ExitIntexpr(c *IntexprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}

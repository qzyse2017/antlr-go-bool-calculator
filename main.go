package main

import (
	"fmt"

	"calculator/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseCalListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	a := ctx.GetText()
	fmt.Println(a)
}

func main() {
	input := antlr.NewInputStream("1<10 AND 1<3")
	lexer := parser.NewCalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Userexpr()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

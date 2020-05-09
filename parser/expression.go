package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	return p.parseExpressionInfix()
}

func (p *Parser) parseExpressionInfix() ast.ExpressionInfix {
	switch p.buffer.Type {
	default:
		return &ast.ExpressionInfixInteger{p.buffer.Value}
	}
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	return p.parseExpressionPrefix()
}

func (p *Parser) parseExpressionPrefix() ast.ExpressionPrefix {
	switch p.buffer.Type {
	default:
		return &ast.ExpressionPrefixInteger{p.buffer.Value}
	}
}

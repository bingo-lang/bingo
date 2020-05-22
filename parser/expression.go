package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	switch p.buffer.Type {
	case token.INTEGER:
		return p.parseExpressionInteger()
	case token.MINUS:
		return p.parseExpressionPrefix()
	default:
		return nil
	}
}

func (p *Parser) parseExpressionInteger() *ast.ExpressionInteger {
	return &ast.ExpressionInteger{p.buffer.Value}
}

func (p *Parser) parseExpressionPrefix() *ast.ExpressionPrefix {
	expressionPrefix := &ast.ExpressionPrefix{Operator: p.buffer}
	p.advance()
	expressionPrefix.Expression = p.parseExpression(LOWEST)
	return expressionPrefix
}

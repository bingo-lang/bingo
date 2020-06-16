package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) parseExpressionUnary() ast.Expression {
	switch p.buffer.Type {
	case token.INTEGER:
		return p.parseExpressionInteger()
	case token.PLUS, token.MINUS:
		return p.parseExpressionPrefix()
	default:
		return nil
	}
}

func (p *Parser) parseExpressionInteger() *ast.ExpressionInteger {
	value := p.buffer.Value
	p.advance()
	return &ast.ExpressionInteger{Value: value}
}

func (p *Parser) parseExpressionPrefix() *ast.ExpressionPrefix {
	expressionPrefix := &ast.ExpressionPrefix{Operator: p.buffer}
	p.advance()
	expressionPrefix.Expression = p.parseExpression(LOWEST)
	return expressionPrefix
}

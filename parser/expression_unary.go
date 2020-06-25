package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) parseExpressionUnary() ast.Expression {
	switch p.token.Type {
	case token.INTEGER:
		return p.parseExpressionInteger()
	case token.PLUS, token.MINUS:
		return p.parseExpressionPrefix()
	default:
		// TODO(tugorez): Handle possible errors.
		return nil
	}
}

func (p *Parser) parseExpressionInteger() *ast.ExpressionInteger {
	value := p.token.Value
	p.advance()
	return &ast.ExpressionInteger{Value: value}
}

func (p *Parser) parseExpressionPrefix() *ast.ExpressionUnary {
	expressionPrefix := &ast.ExpressionUnary{Operator: p.token}
	p.advance()
	expressionPrefix.Expression = p.parseExpression(LOWEST)
	return expressionPrefix
}

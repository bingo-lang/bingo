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
		return nil
	}
}

func (p *Parser) parseExpressionInteger() *ast.ExpressionInteger {
	value := p.token.Value
	p.advance()
	return &ast.ExpressionInteger{Value: value}
}

func (p *Parser) parseExpressionPrefix() *ast.ExpressionUnary {
	prefix := &ast.ExpressionUnary{Operator: p.token}
	p.advance()
	if expression := p.parseExpression(LOWEST); !isNil(expression) {
		prefix.Expression = expression
		return prefix
	}
	return nil
}

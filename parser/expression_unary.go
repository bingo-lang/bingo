package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) parseExpressionUnary() ast.Expression {
	switch p.token.Type {
	case token.PLUS, token.MINUS, token.BANG:
		return p.parseExpressionPrefix()
	case token.INTEGER:
		return p.parseExpressionInteger()
	case token.BOOLEAN:
		return p.parseExpressionBoolean()
	default:
		return nil
	}
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

func (p *Parser) parseExpressionInteger() *ast.ExpressionInteger {
	value := p.token.Value
	p.advance()
	return &ast.ExpressionInteger{Value: value}
}

func (p *Parser) parseExpressionBoolean() *ast.ExpressionBoolean {
	value := p.token.Value
	p.advance()
	return &ast.ExpressionBoolean{Value: value}
}

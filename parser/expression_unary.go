package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) parseExpressionUnary() ast.Expression {
	switch p.token.Type {
	case token.INTEGER:
		return p.parseExpressionInteger()
	case token.BOOLEAN:
		return p.parseExpressionBoolean()
	default:
		return p.parseExpressionPrefix()
	}
}

func (p *Parser) parseExpressionPrefix() ast.ExpressionUnary {
	if p.tokenIsUnaryOperator() {
		return ast.ExpressionUnary{
			Operator:   p.advance(),
			Expression: p.parseExpression(LOWEST),
		}
	}
	// Invalid operator.
	return ast.ExpressionUnary{Operator: p.token, Invalid: true}

}

func (p *Parser) parseExpressionInteger() ast.ExpressionInteger {
	value := p.token.Value
	p.advance()
	return ast.ExpressionInteger{Value: value}
}

func (p *Parser) parseExpressionBoolean() ast.ExpressionBoolean {
	value := p.token.Value
	p.advance()
	return ast.ExpressionBoolean{Value: value}
}

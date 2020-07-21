package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
)

func (p *Parser) parseExpressionUnary() (ast.Expression, error) {
	switch p.token.Type {
	case token.INTEGER:
		return p.parseExpressionInteger()
	case token.BOOLEAN:
		return p.parseExpressionBoolean()
	case token.LPAREN:
		return p.parseExpressionGrouped()
	default:
		return p.parseExpressionPrefix()
	}
}

func (p *Parser) parseExpressionPrefix() (expressionUnary ast.ExpressionUnary, err error) {
	if p.tokenIsUnaryOperator() {
		operator := p.advance()
		if expression, err := p.parseExpression(LOWEST); err == nil {
			expressionUnary = ast.ExpressionUnary{
				Operator:   operator,
				Expression: expression,
			}
		}
	} else {
		err = fmt.Errorf("Token %q is not a unary operator", p.token.Value)
	}
	return
}

func (p *Parser) parseExpressionGrouped() (ast.Expression, error) {
	p.advance()
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	if p.tokenIs(token.RPAREN) {
		p.advance()
		return expression, nil
	}
	return nil, fmt.Errorf("Expecting token to be ) got %s", p.token.Type)
}

func (p *Parser) parseExpressionInteger() (expression ast.ExpressionInteger, err error) {
	value := p.token.Value
	p.advance()
	expression = ast.ExpressionInteger{Value: value}
	return
}

func (p *Parser) parseExpressionBoolean() (expression ast.ExpressionBoolean, err error) {
	value := p.token.Value
	p.advance()
	expression = ast.ExpressionBoolean{Value: value}
	return
}

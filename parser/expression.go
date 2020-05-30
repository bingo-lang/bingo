package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	var expression ast.Expression
	expression = p.parseExpressionUnary()
	for pr := p.precedence(); precedence < pr; pr = p.precedence() {
		infix := &ast.ExpressionInfix{
			ExpressionLeft: expression,
			Operator:       p.buffer,
		}
		p.advance()
		infix.ExpressionRight = p.parseExpression(pr)
		expression = infix
	}
	return expression
}

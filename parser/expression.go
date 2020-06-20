package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	expression := p.parseExpressionUnary()
	// TODO(tugorez): Handle possible errors.
	for pr := p.precedence(); precedence < pr; pr = p.precedence() {
		expression = p.parseExpressionBinary(expression, pr)
	}
	return expression
}

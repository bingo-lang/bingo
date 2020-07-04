package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) (expression ast.Expression) {
	if expression = p.parseExpressionUnary(); isNil(expression) {
		return nil
	}
	for pr := p.precedence(); pr > precedence; pr = p.precedence() {
		if expression = p.parseExpressionBinary(expression, pr); isNil(expression) {
			return nil
		}
	}
	return expression
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) (expression ast.Expression) {
	expression = p.parseExpressionUnary()
	for pr := p.precedence(); pr > precedence && !expression.HasErrors(); pr = p.precedence() {
		expression = p.parseExpressionBinary(expression, pr)
	}
	return expression
}

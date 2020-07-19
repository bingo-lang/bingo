package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) (expression ast.Expression, err error) {
	expression, err = p.parseExpressionUnary()
	for pr := p.precedence(); pr > precedence && err == nil; pr = p.precedence() {
		expression, err = p.parseExpressionBinary(expression, pr)
	}
	return
}

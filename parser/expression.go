package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpression(precedence Precedence) (expression ast.Expression, err error) {
	switch {
	case p.tokenIsInteger():
		expression, err = p.parseExpressionInteger()
	case p.tokenIsBoolean():
		expression, err = p.parseExpressionBoolean()
	case p.tokenIsLParen():
		expression, err = p.parseExpressionGrouped()
	case p.tokenIsUnaryOperator():
		expression, err = p.parseExpressionUnary()
	default:
		err = fmt.Errorf("[ParseExpression] invalid token %q", p.token.Value)
	}
	if err != nil {
		return
	}
	// TODO(tugorez): Create a util function to better validate this precedence validation.
	for pr := p.precedence(); pr > precedence && err == nil; pr = p.precedence() {
		expression, err = p.parseExpressionBinary(expression, pr)
	}
	return
}

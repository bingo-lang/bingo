package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionUnary() (expressionUnary ast.ExpressionUnary, err error) {
	// TODO(tugorez): Validate this is a unary operator.
	operator := p.advance()
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return
	}
	expressionUnary = ast.ExpressionUnary{
		Operator:   operator,
		Expression: expression,
	}
	return
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionUnary() (expressionUnary ast.ExpressionUnary, err error) {
	operator, err := p.assertTokenIsUnaryOperator()
	if err != nil {
		return
	}
	expression, err := p.parseExpression(PREFIX)
	if err != nil {
		return
	}
	expressionUnary = ast.ExpressionUnary{
		Operator:   operator,
		Expression: expression,
	}
	return
}

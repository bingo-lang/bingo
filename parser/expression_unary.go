package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionUnary() (expressionUnary ast.ExpressionUnary, err error) {
	operator := p.token
	if !p.assertTokenIsUnaryOperator() {
		return ast.ExpressionUnary{}, fmt.Errorf("[ExpressionUnary] invalid token %q", operator.Value)
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

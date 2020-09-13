package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) (expression ast.ExpressionBinary, err error) {
	operator, err := p.assertTokenIsBinaryOperator()
	if err != nil {
		return
	}
	expRight, err := p.parseExpression(pr)
	if err != nil {
		return
	}
	expression = ast.ExpressionBinary{
		ExpressionLeft:  expLeft,
		Operator:        operator,
		ExpressionRight: expRight,
	}
	return
}

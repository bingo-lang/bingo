package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) (ast.ExpressionBinary, error) {
	operator := p.token
	if !p.assertTokenIsBinaryOperator() {
		return ast.ExpressionBinary{}, fmt.Errorf("[ExpressionBinary] invalid token %q", operator.Value)
	}
	expRight, err := p.parseExpression(pr)
	if err != nil {
		return ast.ExpressionBinary{}, err
	}
	return ast.ExpressionBinary{
		ExpressionLeft:  expLeft,
		Operator:        operator,
		ExpressionRight: expRight,
	}, nil
}

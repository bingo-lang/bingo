package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) (expressionBinary ast.ExpressionBinary, err error) {
	// TODO(tugorez): Validate this is a binary operator.
	operator := p.advance()
	if expRight, err := p.parseExpression(pr); err == nil {
		expressionBinary = ast.ExpressionBinary{
			ExpressionLeft:  expLeft,
			Operator:        operator,
			ExpressionRight: expRight,
		}
	}
	return
}

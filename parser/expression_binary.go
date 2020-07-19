package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) (expressionBinary ast.ExpressionBinary, err error) {
	// It seems that we can assume the current token is a valid binary operator
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

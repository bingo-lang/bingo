package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) *ast.ExpressionBinary {
	expression := &ast.ExpressionBinary{
		ExpressionLeft: expLeft,
		Operator:       p.token,
	}
	p.advance()
	if expressionRight := p.parseExpression(pr); !isNil(expressionRight) {
		expression.ExpressionRight = expressionRight
		return expression
	}
	return nil
}

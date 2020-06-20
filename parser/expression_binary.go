package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) ast.Expression {
	expression := &ast.ExpressionBinary{
		ExpressionLeft: expLeft,
		Operator:       p.buffer,
	}
	p.advance()
	expression.ExpressionRight = p.parseExpression(pr)
	// TODO(tugorez): Handle possible errors.
	return expression
}

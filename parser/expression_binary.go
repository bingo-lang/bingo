package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBinary(expLeft ast.Expression, pr Precedence) ast.ExpressionBinary {
	return ast.ExpressionBinary{
		ExpressionLeft:  expLeft,
		Operator:        p.advance(),
		ExpressionRight: p.parseExpression(pr),
	}
}

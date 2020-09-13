package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionInteger() (expression ast.ExpressionInteger, err error) {
	integer, err := p.assertTokenIsInteger()
	if err != nil {
		return
	}
	expression = ast.ExpressionInteger{Value: integer.Value}
	return
}

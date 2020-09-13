package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBoolean() (expression ast.ExpressionBoolean, err error) {
	boolean, err := p.assertTokenIsBoolean()
	if err != nil {
		return
	}
	expression = ast.ExpressionBoolean{Value: boolean.Value}
	return
}

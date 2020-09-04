package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBoolean() (expression ast.ExpressionBoolean, err error) {
	value := p.token.Value
	if !p.assertTokenIsBoolean() {
		err = fmt.Errorf("[ExpressionBoolean] invalid token %q", value)
		return
	}
	expression = ast.ExpressionBoolean{Value: value}
	return
}

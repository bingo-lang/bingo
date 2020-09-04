package parser

import (
	"fmt"

	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionInteger() (expression ast.ExpressionInteger, err error) {
	value := p.token.Value
	if !p.assertTokenIsInteger() {
		err = fmt.Errorf("[ExpressionInteger] invalid token %q", value)
		return
	}
	expression = ast.ExpressionInteger{Value: value}
	return
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionBoolean() (expression ast.ExpressionBoolean, err error) {
	value := p.token.Value
	p.advance()
	expression = ast.ExpressionBoolean{Value: value}
	return
}

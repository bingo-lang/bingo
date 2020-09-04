package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionInteger() (expression ast.ExpressionInteger, err error) {
	// TODO(tugorez): Validate it is an integer.
	value := p.token.Value
	p.advance()
	expression = ast.ExpressionInteger{Value: value}
	return
}

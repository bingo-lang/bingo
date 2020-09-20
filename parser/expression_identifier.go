package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionIdentifier() (expression ast.ExpressionIdentifier, err error) {
	identifier, err := p.assertTokenIsIdentifier()
	if err != nil {
		return
	}
	expression = ast.ExpressionIdentifier{Value: identifier.Value}
	return
}

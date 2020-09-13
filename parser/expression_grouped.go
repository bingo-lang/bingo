package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionGrouped() (expression ast.Expression, err error) {
	_, err = p.assertTokenIsLParen()
	if err != nil {
		return
	}
	expression, err = p.parseExpression(LOWEST)
	if err != nil {
		return
	}
	_, err = p.assertTokenIsRParen()
	return
}

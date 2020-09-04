package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseExpressionGrouped() (ast.Expression, error) {
	if !p.assertTokenIsLParen() {
		return nil, fmt.Errorf("[ParseExpressionGrouped] invalid token %q", p.token.Value)
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	if !p.assertTokenIsRParen() {
		return nil, fmt.Errorf("[ParseExpressionGrouped] invalid token %q", p.token.Value)
	}
	return expression, nil
}

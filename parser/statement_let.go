package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatementLet() (statement ast.StatementLet, err error) {
	_, err = p.assertTokenIsLet()
	if err != nil {
		return
	}
	identifier, err := p.assertTokenIsIdentifier()
	if err != nil {
		return
	}
	_, err = p.assertTokenIsAssign()
	if err != nil {
		return
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return
	}
	_, err = p.assertTokenIsSemicolon()
	if err != nil {
		return
	}
	statement = ast.StatementLet{
		Identifier: identifier,
		Expression: expression,
	}
	return
}

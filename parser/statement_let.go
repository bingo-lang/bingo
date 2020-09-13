package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatementLet() (statement ast.StatementLet, err error) {
	_, err = p.assertTokenIsLet()
	if err != nil {
		p.checkIsStatementSeparator()
		return
	}
	identifier, err := p.assertTokenIsIdentifier()
	if err != nil {
		p.checkIsStatementSeparator()
		return
	}
	_, err = p.assertTokenIsAssign()
	if err != nil {
		p.checkIsStatementSeparator()
		return
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		p.checkIsStatementSeparator()
		return ast.StatementLet{}, err
	}
	if token := p.token; !p.checkIsStatementSeparator() {
		err = fmt.Errorf("[StatementLet] invalid token %q", token.Value)
		return
	}
	statement = ast.StatementLet{
		Identifier: identifier,
		Expression: expression,
	}
	return
}

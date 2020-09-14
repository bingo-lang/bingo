package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatementExpression() (statement ast.StatementExpression, err error) {
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return
	}
	_, err = p.assertTokenIsSemicolon()
	if err != nil {
		return
	}
	statement = ast.StatementExpression{Expression: expression}
	return
}

package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	return p.parseStatementExpression()
}

func (p *Parser) parseStatementExpression() (statement ast.StatementExpression, err error) {
	if expression, expressionError := p.parseExpression(LOWEST); expressionError == nil {
		token := p.token
		if p.checkIsStatementSeparator() {
			statement = ast.StatementExpression{Expression: expression}
		} else {
			err = fmt.Errorf("Expecting statatement separator got %q", token)
		}
	} else {
		err = expressionError
		p.checkIsStatementSeparator()
	}
	return
}

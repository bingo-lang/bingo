package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatementExpression() (statement ast.StatementExpression, err error) {
	if expression, expressionError := p.parseExpression(LOWEST); expressionError == nil {
		token := p.token
		if p.checkIsStatementSeparator() {
			statement = ast.StatementExpression{Expression: expression}
		} else {
			err = fmt.Errorf("[StatementExpression] invalid token %q", token.Value)
		}
	} else {
		err = expressionError
		p.checkIsStatementSeparator()
	}
	return
}

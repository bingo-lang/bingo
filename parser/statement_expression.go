package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatementExpression() (statement ast.StatementExpression, err error) {
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		p.checkIsStatementSeparator()
		return
	}
	if token := p.token; !p.checkIsStatementSeparator() {
		err = fmt.Errorf("[StatementExpression] invalid token %q", token.Value)
		return
	}
	return ast.StatementExpression{Expression: expression}, nil
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() (statement ast.Statement, err error) {
	switch {
	case p.tokenIsLBrace():
		statement, err = p.parseStatementBlock()
	case p.tokenIsLet():
		statement, err = p.parseStatementLet()
	default:
		statement, err = p.parseStatementExpression()
	}
	if err != nil {
		p.advanceUntilStatementSeparator()
	}
	return
}

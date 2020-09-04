package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	return p.parseStatementExpression()
}

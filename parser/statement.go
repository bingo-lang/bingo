package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch {
	case p.tokenIsLBrace():
		return p.parseStatementBlock()
	case p.tokenIsLet():
		return p.parseStatementLet()
	default:
		return p.parseStatementExpression()
	}
}

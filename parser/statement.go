package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() ast.Statement {
	switch {
	default:
		return p.parseStatementExpression()
	}
}

func (p *Parser) parseStatementExpression() *ast.StatementExpression {
	expression := p.parseExpression(LOWEST)
	return &ast.StatementExpression{expression}
}

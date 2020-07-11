package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() ast.Statement {
	return p.parseStatementExpression()
}

func (p *Parser) parseStatementExpression() ast.StatementExpression {
	return ast.StatementExpression{
		Expression: p.parseExpression(LOWEST),
		Invalid:    !p.checkIsStatementSeparator(),
	}
}

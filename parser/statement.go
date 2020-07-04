package parser

import (
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatement() ast.Statement {
	if statement := p.parseStatementExpression(); !isNil(statement) && p.tokenIsStatementSeparator() {
		p.advance()
		return statement
	}
	return p.parseStatementError()
}

func (p *Parser) parseStatementExpression() *ast.StatementExpression {
	if expression := p.parseExpression(LOWEST); !isNil(expression) {
		return &ast.StatementExpression{Expression: expression}
	}
	return nil
}

func (p *Parser) parseStatementError() *ast.StatementError {
	statement := &ast.StatementError{InvalidToken: p.token}
	for ; !p.tokenIsStatementSeparator(); p.advance() {
	}
	p.advance()
	return statement
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"go.uber.org/multierr"
)

func (p *Parser) parseStatementBlock() (statement ast.StatementBlock, err error) {
	statements := make([]ast.Statement, 0)
	_, err = p.assertTokenIsLBrace()
	if err != nil {
		return
	}
	for !p.tokenIsEOF() && !p.tokenIsRBrace() {
		if statement, statementError := p.parseStatement(); statementError == nil {
			statements = append(statements, statement)
		} else {
			err = multierr.Combine(err, statementError)
		}
	}
	if err != nil {
		return ast.StatementBlock{}, err
	}
	_, err = p.assertTokenIsRBrace()
	if err != nil {
		return
	}
	statement = ast.StatementBlock{Statements: statements}
	return
}

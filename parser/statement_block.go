package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"go.uber.org/multierr"
)

func (p *Parser) parseStatementBlock() (statement ast.StatementBlock, err error) {
	_, err = p.assertTokenIsLBrace()
	if err != nil {
		return
	}
	statements, err := p.parseStatements()
	if err != nil {
		return
	}
	_, err = p.assertTokenIsRBrace()
	if err != nil {
		return
	}
	statement = ast.StatementBlock{Statements: statements}
	return
}

func (p *Parser) parseStatements() (statements []ast.Statement, err error) {
	statements = make([]ast.Statement, 0)
	for !p.tokenIsEOF() && !p.tokenIsRBrace() {
		if stmt, stmtError := p.parseStatement(); stmtError == nil {
			statements = append(statements, stmt)
		} else {
			err = multierr.Combine(err, stmtError)
		}
	}
	return
}

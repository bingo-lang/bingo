package parser

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
)

func (p *Parser) parseStatementLet() (statement ast.StatementLet, err error) {
	let := p.token
	if !p.assertTokenIsLet() {
		return ast.StatementLet{}, fmt.Errorf("[StatementLet] invalid token %q", let.Value)
	}
	identifier := p.token
	if !p.assertTokenIsIdentifier() {
		return ast.StatementLet{}, fmt.Errorf("[StatementLet] invalid token %q", identifier.Value)
	}
	assign := p.token
	if !p.assertTokenIsAssign() {
		return ast.StatementLet{}, fmt.Errorf("[StatementLet] invalid token %q", assign.Value)
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return ast.StatementLet{}, err
	}
	statement = ast.StatementLet{
		Identifier: identifier,
		Expression: expression,
	}
	return
}
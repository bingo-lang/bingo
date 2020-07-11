package parser

import (
	"github.com/bingo-lang/bingo/token"
	"reflect"
)

func (p *Parser) tokenIsUnaryOperator() bool {
	switch p.token.Type {
	case token.PLUS, token.MINUS,
		token.BANG:
		return true
	default:
		return false
	}
}

func (p *Parser) checkIsStatementSeparator() bool {
	if p.tokenIsStatementSeparator() {
		p.advance()
		return true
	}
	for ; !p.tokenIsStatementSeparator(); p.advance() {
	}
	return false
}

func (p *Parser) tokenIsEOF() bool {
	return p.tokenIs(token.EOF)
}

func (p *Parser) tokenIsStatementSeparator() bool {
	return p.tokenIsEOF() || p.tokenIs(token.SEMICOLON)
}

func (p *Parser) tokenIs(typ token.Type) bool {
	return p.token.Type == typ
}

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

package parser

import (
	"github.com/bingo-lang/bingo/token"
	"reflect"
)

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

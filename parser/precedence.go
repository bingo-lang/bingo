package parser

import (
	"github.com/bingo-lang/bingo/token"
)

type Precedence int

const (
	LOWEST = iota
	LOGICAL
	ADDITION
	MULTIPLICATION
)

func (p *Parser) precedence() Precedence {
	switch p.token.Type {
	case token.PLUS, token.MINUS:
		return ADDITION
	case token.ASTERISK, token.SLASH:
		return MULTIPLICATION
	case token.OR, token.AND, token.GT, token.GTE, token.LT, token.LTE, token.EQUAL:
		return LOGICAL
	default:
		return LOWEST
	}
}

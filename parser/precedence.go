package parser

import (
	"github.com/bingo-lang/bingo/token"
)

type Precedence int

const (
	LOWEST = iota
	AND_OR_OR
	LOGICAL
	ADDITION
	MULTIPLICATION
)

func (p *Parser) precedence() Precedence {
	switch p.token.Type {
	case token.AND, token.OR:
		return AND_OR_OR
	case token.GT, token.GTE, token.LT, token.LTE, token.EQUAL:
		return LOGICAL
	case token.PLUS, token.MINUS:
		return ADDITION
	case token.ASTERISK, token.SLASH:
		return MULTIPLICATION
	default:
		return LOWEST
	}
}

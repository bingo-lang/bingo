package parser

import (
	"github.com/bingo-lang/bingo/token"
)

type Precedence int

const (
	LOWEST = iota
	AND
	OR
	GT
	GTE
	LT
	LTE
	EQUAL
	MINUS
	PLUS
	SLASH
	ASTERISK
	// Prefix operators.
	PREFIX
)

func (p *Parser) precedence() Precedence {
	switch p.token.Type {
	case token.OR:
		return OR
	case token.AND:
		return AND

	case token.GT:
		return GT
	case token.LT:
		return LT

	case token.LTE:
		return LTE
	case token.GTE:
		return GTE

	case token.EQUAL:
		return EQUAL

	case token.MINUS:
		return MINUS
	case token.PLUS:
		return PLUS

	case token.SLASH:
		return SLASH
	case token.ASTERISK:
		return ASTERISK

	default:
		return LOWEST
	}
}

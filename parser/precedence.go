package parser

import (
	"github.com/bingo-lang/bingo/token"
)

type Precedence int

const (
	LOWEST = iota
	ADDITION
	MULTIPLICATION
)

func (p *Parser) precedence() Precedence {
	switch p.buffer.Type {
	case token.PLUS, token.MINUS:
		return ADDITION
	case token.ASTERISK, token.SLASH:
		return MULTIPLICATION
	default:
		return LOWEST
	}
}

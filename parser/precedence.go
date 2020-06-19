package parser

import (
	"github.com/bingo-lang/bingo/token"
)

type Precedence int

const (
	LOWEST = iota
	ADDITION
)

func (p *Parser) precedence() Precedence {
	switch p.buffer.Type {
	case token.PLUS, token.MINUS:
		return ADDITION
	default:
		return LOWEST
	}
}

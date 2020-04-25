package scanner

import (
	"github.com/bingo-lang/bingo/token"
)

func (s *Scanner) nextToken() token.Token {
	switch {
	case isDigit(s.buffer):
		return s.nextNumber()
	default:
		return token.New(token.UNDEFINED, "")
	}
}

func (s *Scanner) nextNumber() token.Token {
	number := ""
	for ; isDigit(s.buffer); s.advance() {
		number += string(s.buffer)
	}
	return token.New(token.INTEGER, number)
}

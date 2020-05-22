package scanner

import (
	"github.com/bingo-lang/bingo/token"
)

func (s *Scanner) nextToken() token.Token {
	switch {
	case s.eof:
		return s.nextEof()
	case isDigit(s.buffer):
		return s.nextNumber()
	default:
		return s.nextUndefined()
	}
}

func (s *Scanner) nextNumber() token.Token {
	number := ""
	for ; isDigit(s.buffer); s.advance() {
		number += string(s.buffer)
	}
	return token.New(token.INTEGER, number)
}

func (s *Scanner) nextEof() token.Token {
	return token.New(token.EOF, "")
}

func (s *Scanner) nextUndefined() token.Token {
	undefined := string(s.buffer)
	s.advance()
	return token.New(token.UNDEFINED, undefined)
}

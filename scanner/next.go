package scanner

import (
	"github.com/bingo-lang/bingo/token"
)

func (s *Scanner) nextToken() token.Token {
	switch {
	case s.eof:
		return s.nextEof()
	case isDigit(s.char):
		return s.nextNumber()
	case isSymbol(s.char):
		return s.nextSymbol()
	default:
		return s.nextUndefined()
	}
}

func (s *Scanner) nextEof() token.Token {
	return token.New(token.EOF, "")
}

func (s *Scanner) nextNumber() token.Token {
	number := ""
	for ; isDigit(s.char); s.advance() {
		number += string(s.char)
	}
	return token.New(token.INTEGER, number)
}

func (s *Scanner) nextSymbol() token.Token {
	switch s.char {
	case '+':
		symbol := string(s.char)
		s.advance()
		return token.New(token.PLUS, symbol)
	case '-':
		symbol := string(s.char)
		s.advance()
		return token.New(token.MINUS, symbol)
	case '*':
		symbol := string(s.char)
		s.advance()
		return token.New(token.ASTERISK, symbol)
	case '/':
		symbol := string(s.char)
		s.advance()
		return token.New(token.SLASH, symbol)
	default:
		return s.nextUndefined()
	}
}

func (s *Scanner) nextUndefined() token.Token {
	undefined := string(s.char)
	s.advance()
	return token.New(token.UNDEFINED, undefined)
}

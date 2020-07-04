package scanner

import (
	"github.com/bingo-lang/bingo/token"
)

func (s *Scanner) scanToken() token.Token {
	switch {
	case s.eof:
		return s.scanEof()
	case isDigit(s.char):
		return s.scanNumber()
	case isSymbol(s.char):
		return s.scanSymbol()
	case isLetter(s.char):
		return s.scanWord()
	default:
		return s.scanUndefined()
	}
}

func (s *Scanner) scanEof() token.Token {
	return token.New(token.EOF, "")
}

func (s *Scanner) scanNumber() token.Token {
	number := ""
	for ; isDigit(s.char); s.advance() {
		number += string(s.char)
	}
	return token.New(token.INTEGER, number)
}

func (s *Scanner) scanSymbol() token.Token {
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
	case '=':
		symbol := string(s.char)
		s.advance()
		return token.New(token.EQUALS, symbol)
	case ';':
		symbol := string(s.char)
		s.advance()
		return token.New(token.SEMICOLON, symbol)
	default:
		return s.scanUndefined()
	}
}

func (s *Scanner) scanWord() token.Token {
	word := ""
	for ; isLetter(s.char); s.advance() {
		word += string(s.char)
	}
	switch word {
	case "let":
		return token.New(token.LET, word)
	case "int":
		return token.New(token.INT, word)
	default:
		return token.New(token.IDENTIFIER, word)
	}
}

func (s *Scanner) scanUndefined() token.Token {
	undefined := string(s.char)
	s.advance()
	return token.New(token.UNDEFINED, undefined)
}

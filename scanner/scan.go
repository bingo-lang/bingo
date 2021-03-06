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
		if s.char == '=' {
			symbol += string(s.char)
			s.advance()
			return token.New(token.EQUAL, symbol)
		}
		return token.New(token.ASSIGN, symbol)
	case ';':
		symbol := string(s.char)
		s.advance()
		return token.New(token.SEMICOLON, symbol)
	case '|':
		symbol := string(s.char)
		s.advance()
		if s.char == '|' {
			symbol += string(s.char)
			s.advance()
			return token.New(token.OR, symbol)
		} else {
			return s.scanUndefined()
		}
	case '&':
		symbol := string(s.char)
		s.advance()
		if s.char == '&' {
			symbol += string(s.char)
			s.advance()
			return token.New(token.AND, symbol)
		} else {
			return s.scanUndefined()
		}
	case '>':
		symbol := string(s.char)
		s.advance()
		if s.char == '=' {
			symbol += string(s.char)
			s.advance()
			return token.New(token.GTE, symbol)
		}
		return token.New(token.GT, symbol)
	case '<':
		symbol := string(s.char)
		s.advance()
		if s.char == '=' {
			symbol += string(s.char)
			s.advance()
			return token.New(token.LTE, symbol)
		}
		return token.New(token.LT, symbol)
	case '!':
		symbol := string(s.char)
		s.advance()
		return token.New(token.BANG, symbol)
	case '(':
		symbol := string(s.char)
		s.advance()
		return token.New(token.LPAREN, symbol)
	case ')':
		symbol := string(s.char)
		s.advance()
		return token.New(token.RPAREN, symbol)
	case '{':
		symbol := string(s.char)
		s.advance()
		return token.New(token.LBRACE, symbol)
	case '}':
		symbol := string(s.char)
		s.advance()
		return token.New(token.RBRACE, symbol)

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
	case "bool":
		return token.New(token.BOOL, word)
	case "true", "false":
		return token.New(token.BOOLEAN, word)
	default:
		return token.New(token.IDENTIFIER, word)
	}
}

func (s *Scanner) scanUndefined() token.Token {
	undefined := string(s.char)
	s.advance()
	return token.New(token.UNDEFINED, undefined)
}

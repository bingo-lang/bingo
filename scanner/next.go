package scanner

import (
	"github.com/bingo-lang/bingo/token"
)

func (s *Scanner) nextToken() token.Token {
	if isLetter(s.buffer) {
		return s.nextWord()
	}
	if isDigit(s.buffer) {
		return s.nextNumber()
	}
	if isSymbol(s.buffer) {
		return s.nextSymbol()
	}
	return token.Token{}
}

func (s *Scanner) nextSpace() token.Token {
	space := ""
	for ; isSpace(s.buffer); s.advance() {
	}
	return token.New(token.SPACE, space)
}

func (s *Scanner) nextWord() token.Token {
	word := ""
	for ; isLetter(s.buffer); s.advance() {
		word += string(s.buffer)
	}
	switch word {
	case "let":
		return token.New(token.LET, word)
	default:
		return token.New(token.IDENTIFIER, word)
	}
}

func (s *Scanner) nextNumber() token.Token {
	number := ""
	for ; isDigit(s.buffer); s.advance() {
		number += string(s.buffer)
	}
	return token.New(token.INTEGER, number)
}

func (s *Scanner) nextSymbol() token.Token {
	symbol := string(s.buffer)
	s.advance()
	switch symbol {
	case "=":
		return token.New(token.ASSIGNMENT, symbol)
	default:
		return token.New(token.SEMICOLON, symbol)
	}
}

package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"io"
)

type Scanner struct {
	source io.RuneReader

	buffer rune
	eof    bool
}

func New(source io.RuneReader) (*Scanner, error) {
	scanner := &Scanner{source: source}
	scanner.advance()
	return scanner, nil
}

func (s *Scanner) NextToken() token.Token {
	s.nextSpace()
	token := s.nextToken()
	s.nextSpace()
	return token
}

func (s *Scanner) EOF() bool {
	return s.eof
}

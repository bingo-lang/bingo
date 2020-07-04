package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"io"
)

type Scanner struct {
	source io.RuneReader

	char rune
	eof  bool
}

func New(source io.RuneReader) *Scanner {
	scanner := &Scanner{source: source}
	scanner.advance()
	return scanner
}

func (s *Scanner) ScanToken() token.Token {
	s.removeSpace()
	token := s.scanToken()
	s.removeSpace()
	return token
}

package scanner

import (
	"github.com/bingo-lang/bingo/token"
)

type Scanner struct {
	source string
}

func New(source string) (*Scanner, error) {
	return &Scanner{source}, nil
}

func (s *Scanner) Token() token.Token {
	return token.Token{}
}

func (s *Scanner) IsDone() bool {
	return true
}

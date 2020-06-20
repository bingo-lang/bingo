package scanner

func (s *Scanner) advance() {
	r, _, err := s.source.ReadRune()
	if err != nil {
		s.eof = true
	}
	s.char = r
}

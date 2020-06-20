package scanner

func (s *Scanner) removeSpace() {
	for ; isSpace(s.char); s.advance() {
	}
}

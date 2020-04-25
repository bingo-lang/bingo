package scanner

func (s *Scanner) removeSpace() {
	for ; isSpace(s.buffer); s.advance() {
	}
}

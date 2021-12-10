package main

type ChunkRuneStack []ChunkRune

func (s *ChunkRuneStack) Push(cr ChunkRune) {
	if s == nil {
		*s = ChunkRuneStack{cr}
	} else {
		*s = append(*s, cr)
	}
}

func (s *ChunkRuneStack) Pop() (ChunkRune, bool) {
	cr, ok := s.Peek()
	if ok {
		*s = (*s)[:len(*s)-1]
	}
	return cr, ok
}

func (s *ChunkRuneStack) Peek() (ChunkRune, bool) {
	if s == nil || len(*s) == 0 {
		return ChunkRune{}, false
	}
	return (*s)[len(*s)-1], true
}

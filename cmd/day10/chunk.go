package main

import "strings"

type Chunk string

const (
	ChunkInvalid Chunk = ""
	Chunk0       Chunk = "()"
	Chunk1       Chunk = "[]"
	Chunk2       Chunk = "{}"
	Chunk3       Chunk = "<>"
)

func (c Chunk) Open() string {
	return string(c)[0:1]
}

func (c Chunk) Close() string {
	return string(c)[1:]
}

func (c Chunk) SyntaxErrScore() int {
	switch c {
	case Chunk0:
		return 3
	case Chunk1:
		return 57
	case Chunk2:
		return 1197
	case Chunk3:
		return 25137
	default:
		return 0
	}
}

func (c Chunk) AutocompleteScore() int {
	switch c {
	case Chunk0:
		return 1
	case Chunk1:
		return 2
	case Chunk2:
		return 3
	case Chunk3:
		return 4
	default:
		return 0
	}
}

type ChunkRune struct {
	Chunk
	open bool
}

func (cr ChunkRune) String() string {
	if cr.open {
		return cr.Open()
	}
	return cr.Close()
}

func ParseChunkRune(r rune) (ChunkRune, bool) {
	open := strings.ContainsRune("([{<", r)
	switch r {
	case '(', ')':
		return ChunkRune{Chunk0, open}, true
	case '[', ']':
		return ChunkRune{Chunk1, open}, true
	case '{', '}':
		return ChunkRune{Chunk2, open}, true
	case '<', '>':
		return ChunkRune{Chunk3, open}, true
	}
	return ChunkRune{}, false
}

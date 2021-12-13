package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func newPathScanner(reader io.Reader) *pathScanner {
	return &pathScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type pathScanner struct {
	scanner *bufio.Scanner
	path    Path
	err     error
}

func (s *pathScanner) Scan() bool {
	if !s.scanner.Scan() {
		s.err = s.scanner.Err()
		return false
	}
	text := s.scanner.Text()
	sep := strings.IndexByte(text, '-')
	if sep == -1 {
		s.err = fmt.Errorf("missing delimiter '-' from line: %q", text)
		return false
	}
	s.path = Path{Cave(text[:sep]), Cave(text[sep+1:])}
	return true
}

func (s *pathScanner) Path() Path {
	return s.path
}

func (s *pathScanner) Err() error {
	return s.err
}

package main

import (
	"bufio"
	"fmt"
	"io"
)

func newEntryScanner(reader io.Reader) *entryScanner {
	return &entryScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type entryScanner struct {
	scanner *bufio.Scanner
	err     error
	entry   Entry
}

func (s *entryScanner) Scan() bool {
	if !s.scanner.Scan() {
		s.err = s.scanner.Err()
		return false
	}
	var sig [10]string
	var out [4]string
	_, s.err = fmt.Sscanf(s.scanner.Text(),
		"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
		&sig[0], &sig[1], &sig[2], &sig[3], &sig[4],
		&sig[5], &sig[6], &sig[7], &sig[8], &sig[9],
		&out[0], &out[1], &out[2], &out[3],
	)
	if s.err != nil {
		return false
	}
	s.entry = Entry{}
	for i, d := range sig {
		s.entry.SignalPatterns[i] = Digit(d)
	}
	for i, d := range out {
		s.entry.OutputValue[i] = Digit(d)
	}
	return true
}

func (s *entryScanner) Entry() Entry {
	return s.entry
}

func (s *entryScanner) Err() error {
	return s.err
}

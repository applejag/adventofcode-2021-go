package main

import (
	"bufio"
	"io"
)

func NewIntScanner(reader io.Reader) IntScanner {
	return &intScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type intScanner struct {
	scanner *bufio.Scanner
	integer int
	err     error
}

func (is *intScanner) Scan() bool {
	if !is.scanner.Scan() {
		return false
	}
	is.integer, is.err = parseInt(is.scanner.Text())
	return is.err == nil
}

func (is *intScanner) Int() int {
	return is.integer
}

func (is *intScanner) Err() error {
	if is.err != nil {
		return is.err
	}
	return is.scanner.Err()
}

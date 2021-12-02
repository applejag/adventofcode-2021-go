package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func NewCmdScanner(reader io.Reader) *cmdScanner {
	return &cmdScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type cmdScanner struct {
	scanner *bufio.Scanner
	cmd     Command
	arg     int
	err     error
}

func (cs *cmdScanner) Scan() bool {
	if !cs.scanner.Scan() {
		return false
	}
	_, cs.err = fmt.Sscanf(cs.scanner.Text(), "%s %d", &cs.cmd, &cs.arg)
	return cs.err == nil
}

func (cs *cmdScanner) Command() Command {
	return cs.cmd
}

func (cs *cmdScanner) Argument() int {
	return cs.arg
}

func (cs *cmdScanner) Err() error {
	return cs.err
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

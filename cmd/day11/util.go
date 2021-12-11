package main

import (
	"fmt"
)

type Empty struct{}

func parseDigit(d rune) (byte, error) {
	if d < '0' || d > '9' {
		return 0, fmt.Errorf("invalid digit: %s", string(d))
	}
	return byte(d - '0'), nil
}

package util

import (
	"errors"
	"fmt"
)

var ErrInvalidDigit = errors.New("invalid digit rune")

func ParseDigitRune(d rune) (int, error) {
	if d < '0' || d > '9' {
		return 0, fmt.Errorf("%w: %q", ErrInvalidDigit, string(d))
	}
	return int(d - '0'), nil
}

func SumInts(vals []int) int {
	var sum int
	for _, v := range vals {
		sum += v
	}
	return sum
}

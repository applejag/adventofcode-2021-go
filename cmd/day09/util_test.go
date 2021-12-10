package main

import "testing"

func TestParseDigit(t *testing.T) {
	digits := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	for want, d := range digits {
		t.Run(string(d), func(t *testing.T) {
			got, err := parseDigit(d)
			if err != nil {
				t.Error(err)
			}
			if int(got) != want {
				t.Errorf("want %d, but got %d", want, got)
			}
		})
	}
}

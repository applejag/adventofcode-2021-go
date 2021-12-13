package main

import "strings"

type Cave string

const (
	CaveStart Cave = "start"
	CaveEnd   Cave = "end"
)

func (c Cave) IsSmall() bool {
	return c[0] >= 'a' && c[0] <= 'z'
}

type Path []Cave

func (p Path) String() string {
	var sb strings.Builder
	for i, c := range p {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(string(c))
	}
	return sb.String()
}

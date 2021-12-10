package main

import (
	"fmt"
	"strings"
)

type Segment byte

const (
	SegA Segment = 1 << iota
	SegB
	SegC
	SegD
	SegE
	SegF
	SegG
	SegNone Segment = 0
)

func (s Segment) String() string {
	switch s {
	case SegA:
		return "a"
	case SegB:
		return "b"
	case SegC:
		return "c"
	case SegD:
		return "d"
	case SegE:
		return "e"
	case SegF:
		return "f"
	case SegG:
		return "g"
	default:
		return fmt.Sprintf("%T(%d)", s, s)
	}
}

func (s Segment) Count() int {
	var count int
	for i := SegA; i <= SegG; i <<= 1 {
		if s&i != 0 {
			count++
		}
	}
	return count
}

func ParseSegment(r rune) (Segment, error) {
	switch r {
	case 'a':
		return SegA, nil
	case 'b':
		return SegB, nil
	case 'c':
		return SegC, nil
	case 'd':
		return SegD, nil
	case 'e':
		return SegE, nil
	case 'f':
		return SegF, nil
	case 'g':
		return SegG, nil
	default:
		return SegNone, fmt.Errorf("unknown segment: %s", r)
	}
}

func ParseSegments(s string) (Segment, error) {
	var segSum Segment
	for _, r := range s {
		seg, err := ParseSegment(r)
		if err != nil {
			return SegNone, err
		}
		segSum |= seg
	}
	return segSum, nil
}

type Entry struct {
	SignalPatterns [10]Segment
	OutputValue    [4]Segment
}

func (e Entry) String() string {
	var sb strings.Builder
	for _, d := range e.SignalPatterns {
		sb.WriteString(d.String())
		sb.WriteByte(' ')
	}
	sb.WriteByte('|')
	for _, d := range e.OutputValue {
		sb.WriteByte(' ')
		sb.WriteString(d.String())
	}
	return sb.String()
}

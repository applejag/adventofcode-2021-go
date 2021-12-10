package main

import (
	"fmt"
	"strings"
)

type Segments byte

const (
	SegA Segments = 1 << iota
	SegB
	SegC
	SegD
	SegE
	SegF
	SegG
	SegNone Segments = 0
)

const (
	Seg0 Segments = SegA | SegB | SegC | SegE | SegF | SegG
	Seg1 Segments = SegC | SegF
	Seg2 Segments = SegA | SegC | SegD | SegE | SegG
	Seg3 Segments = SegA | SegC | SegD | SegF | SegG
	Seg4 Segments = SegB | SegC | SegD | SegF
	Seg5 Segments = SegA | SegB | SegD | SegF | SegG
	Seg6 Segments = SegA | SegB | SegD | SegE | SegF | SegG
	Seg7 Segments = SegA | SegC | SegF
	Seg8 Segments = SegA | SegB | SegC | SegD | SegE | SegF | SegG
	Seg9 Segments = SegA | SegB | SegC | SegD | SegF | SegG
)

func (s Segments) String() string {
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
		var sb strings.Builder
		for i := SegA; i <= SegG; i <<= 1 {
			singleSeg := s & i
			if singleSeg != SegNone {
				sb.WriteString(singleSeg.String())
			}
		}
		return sb.String()
	}
}

func (s Segments) Count() int {
	var count int
	for i := SegA; i <= SegG; i <<= 1 {
		if s&i != 0 {
			count++
		}
	}
	return count
}

func (s Segments) Int() (int, error) {
	switch s {
	case Seg0:
		return 0, nil
	case Seg1:
		return 1, nil
	case Seg2:
		return 2, nil
	case Seg3:
		return 3, nil
	case Seg4:
		return 4, nil
	case Seg5:
		return 5, nil
	case Seg6:
		return 6, nil
	case Seg7:
		return 7, nil
	case Seg8:
		return 8, nil
	case Seg9:
		return 9, nil
	default:
		return -1, fmt.Errorf("segments does not represent a digit: %s", s)
	}
}

func ParseSegment(r rune) (Segments, error) {
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
		return SegNone, fmt.Errorf("unknown segment: %v", r)
	}
}

func ParseSegments(s string) (Segments, error) {
	var segSum Segments
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
	SignalPatterns [10]Segments
	OutputValue    [4]Segments
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

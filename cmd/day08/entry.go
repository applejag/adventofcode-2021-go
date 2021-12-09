package main

import "strings"

type Segment byte

const (
	SegA Segment = 'a'
	SegB Segment = 'b'
	SegC Segment = 'c'
	SegD Segment = 'd'
	SegE Segment = 'e'
	SegF Segment = 'f'
	SegG Segment = 'g'
)

func (s Segment) String() string {
	return string(byte(s))
}

type Digit []Segment

func (d Digit) String() string {
	bytes := make([]byte, len(d))
	for i, b := range d {
		bytes[i] = byte(b)
	}
	return string(bytes)
}

type Entry struct {
	SignalPatterns [10]Digit
	OutputValue    [4]Digit
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

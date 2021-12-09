package main

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

type Digit []Segment

type Entry struct {
	SignalPatterns [10]Digit
	OutputValue    [4]Digit
}

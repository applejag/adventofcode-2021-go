package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day08")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	var scanner = newEntryScanner(inputFile)
	var entries []Entry

	for scanner.Scan() {
		entries = append(entries, scanner.Entry())
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).
			Message("Failed to scan.")
		os.Exit(1)
	}

	log.Info().WithInt("entries", len(entries)).
		Message("Scanning complete.")

	if common.Part2 {
		part2(entries)
	} else {
		part1(entries)
	}
}

func part1(entries []Entry) {
	var uniqueSegDigits int
	for _, entry := range entries {
		for _, out := range entry.OutputValue {
			switch out.Count() {
			case 2, 3, 4, 7:
				uniqueSegDigits++
			}
		}
	}
	log.Info().WithInt("digits", uniqueSegDigits).
		Message("Counted unique segment-count digits in outputs.")
}

type segEquality struct {
	signal Segments
	target Segments
}

func (se segEquality) String() string {
	return fmt.Sprintf("%s=%s", se.signal, strings.ToUpper(se.target.String()))
}

func part2(entries []Entry) {
	var sum int

	for _, entry := range entries {
		var (
			segEq      []segEquality
			seg5, seg6 []Segments
		)
		for _, signal := range entry.SignalPatterns {
			switch signal.Count() {
			case 2:
				segEq = append(segEq, segEquality{signal, SegC | SegF})
			case 3:
				segEq = append(segEq, segEquality{signal, SegA | SegC | SegF})
			case 4:
				segEq = append(segEq, segEquality{signal, SegB | SegC | SegD | SegF})
			case 5:
				seg5 = append(seg5, signal)
			case 6:
				seg6 = append(seg6, signal)
				// ignore 7 segment digits, i.e. digit 8, as we cannot deduce
				// anything from it
			}
		}
		segEq = append(segEq, deduceEqualities5Seg(seg5)...)
		segEq = append(segEq, deduceEqualities6Seg(seg6)...)
		log.Debug().WithStringf("segEq", "%v", segEq).Message("Deduced equalities")
		m := map[Segments]Segments{
			SegA: intersectAllFor(segEq, SegA),
			SegB: intersectAllFor(segEq, SegB),
			SegC: intersectAllFor(segEq, SegC),
			SegD: intersectAllFor(segEq, SegD),
			SegE: intersectAllFor(segEq, SegE),
			SegF: intersectAllFor(segEq, SegF),
			SegG: intersectAllFor(segEq, SegG),
		}
		var uniqueMatches []Segments
		for _, b := range m {
			if b.Count() == 1 {
				uniqueMatches = append(uniqueMatches, b)
			}
		}
		for a, b := range m {
			if b.Count() <= 1 {
				continue
			}
			for _, unique := range uniqueMatches {
				// subtract known matches
				m[a] &= ^unique
			}
		}
		log.Debug().WithStringf("m", "%v", m).Message("Intersected")
		var output int
		for _, out := range entry.OutputValue {
			val, err := translateSegments(m, out).Int()
			if err != nil {
				log.Error().WithError(err).WithStringer("entry", entry).
					Message("Failed to translate segments")
				os.Exit(1)
			}
			output = output*10 + val
		}
		log.Debug().WithInt("output", output).Message("Converted entry out to int")
		sum += output
	}
	log.Info().WithInt("sum", sum).Message("Summed up all outputs")
}

func translateSegments(m map[Segments]Segments, seg Segments) Segments {
	var res Segments
	for i := SegA; i <= SegG; i <<= 1 {
		singleSeg := seg & i
		if singleSeg != SegNone {
			res |= m[singleSeg]
		}
	}
	return res
}

func intersectAllFor(segEq []segEquality, seg Segments) Segments {
	var segments []Segments
	for _, eq := range segEq {
		if eq.signal&seg != 0 {
			segments = append(segments, eq.target)
		}
	}
	if len(segments) == 0 {
		return SegNone
	}
	res := segments[0]
	for i := 1; i < len(segments); i++ {
		res &= segments[i]
	}
	return res
}

func deduceEqualities5Seg(segs []Segments) []segEquality {
	counts := countSegments(segs)
	var (
		seg3 Segments
		seg2 Segments
		seg1 Segments
	)
	for seg, count := range counts {
		switch count {
		case 3:
			seg3 |= seg
		case 2:
			seg2 |= seg
		case 1:
			seg1 |= seg
		}
	}
	return []segEquality{
		// |  aaa  |
		// |       |
		// |       |
		// |  ddd  |
		// |       |
		// |       |
		// |  ggg  |
		{seg3, SegA | SegD | SegG},
		// |       |
		// |     c |
		// |     c |
		// |       |
		// |     f |
		// |     f |
		// |       |
		{seg2, SegC | SegF},
		// |       |
		// | b     |
		// | b     |
		// |       |
		// | e     |
		// | e     |
		// |       |
		{seg1, SegB | SegE},
	}
}

func deduceEqualities6Seg(segs []Segments) []segEquality {
	counts := countSegments(segs)
	var (
		seg3 Segments
		seg2 Segments
	)
	for seg, count := range counts {
		switch count {
		case 3:
			seg3 |= seg
		case 2:
			seg2 |= seg
		}
	}
	return []segEquality{
		// |  aaa  |
		// | b     |
		// | b     |
		// |       |
		// |     f |
		// |     f |
		// |  ggg  |
		{seg3, SegA | SegB | SegF | SegG},
		// |       |
		// |     c |
		// |     c |
		// |  ddd  |
		// | e     |
		// | e     |
		// |       |
		{seg2, SegC | SegD | SegE},
	}
}

func countSegments(segs []Segments) map[Segments]int {
	counts := map[Segments]int{}
	for _, seg := range segs {
		for i := SegA; i <= SegG; i <<= 1 {
			singleSeg := seg & i
			if singleSeg == SegNone {
				continue
			}
			counts[singleSeg]++
		}
	}
	return counts
}

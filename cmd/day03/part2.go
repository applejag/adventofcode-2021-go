package main

import (
	"fmt"
	"io"
)

type Bits struct {
	Num  int
	Size int
}

func (b Bits) Bit(idx int) int {
	return (b.Num >> (b.Size - idx - 1)) & 1
}

func (b Bits) String() string {
	return fmt.Sprintf(fmt.Sprintf("%%0%db", b.Size), b.Num)
}

func part2(inputFile io.Reader) {
	scanner := NewBitScanner(inputFile)

	var allBits []Bits
	for scanner.Scan() {
		allBits = append(allBits, Bits{
			Num:  scanner.Bits(),
			Size: scanner.BitsCount(),
		})
	}
	log.Info().WithInt("scans", len(allBits)).
		Message("Scanning complete. Now processing life support rating.")

	oxygen := filterFunc(filterMostCommon, allBits)
	co2 := filterFunc(filterLeastCommon, allBits)

	log.Info().WithStringer("oxygen", oxygen).
		WithStringer("co2", co2).
		WithInt("life", oxygen.Num*co2.Num).
		Message("Life support rating.")
}

type BitsFilter func(bitsSlice []Bits, idx int) []Bits

func filterFunc(filter BitsFilter, bitsSlice []Bits) Bits {
	idx := 0
	for len(bitsSlice) > 1 {
		bitsSlice = filter(bitsSlice, idx)
		idx++
	}
	return bitsSlice[0]
}

func filterMostCommon(bitsSlice []Bits, idx int) []Bits {
	mostCommonBit := findCommonBit(bitsSlice, idx)
	log.Debug().
		WithInt("idx", idx).
		WithInt("mostCommon", mostCommonBit).
		Message("Filter most common.  ")
	return filterBits(bitsSlice, idx, mostCommonBit)
}

func filterLeastCommon(bitsSlice []Bits, idx int) []Bits {
	leastCommonBit := 1 - findCommonBit(bitsSlice, idx)
	log.Debug().
		WithInt("idx", idx).
		WithInt("mostCommon", leastCommonBit).
		Message("Filter least common. ")
	return filterBits(bitsSlice, idx, leastCommonBit)
}

func filterBits(bitsSlice []Bits, idx, want int) []Bits {
	var filtered []Bits
	for _, bits := range bitsSlice {
		if bits.Bit(idx) == want {
			filtered = append(filtered, bits)
		}
	}
	return filtered
}

func findCommonBit(bitsSlice []Bits, idx int) int {
	ones := sumBitsAtIndex(bitsSlice, idx)
	zeroes := len(bitsSlice) - ones
	log.Debug().
		WithInt("idx", idx).
		WithInt("sum", ones).
		WithInt("len", len(bitsSlice)).
		Message("Find most common bit.")
	if ones >= zeroes {
		return 1
	} else {
		return 0
	}
}

func sumBitsAtIndex(bitsSlice []Bits, idx int) int {
	var sum int
	for _, bits := range bitsSlice {
		sum += bits.Bit(idx)
	}
	return sum
}

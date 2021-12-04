package main

import (
	"io"
	"os"
)

func part1(inputFile io.Reader) {
	scanner := NewBitScanner(inputFile)
	var bitSums []int
	var scansCount int

	for scanner.Scan() {
		bits := scanner.Bits()
		bitCount := scanner.BitsCount()
		scansCount++
		bitSums = growSliceToLen(bitSums, bitCount)

		addBits(bits, bitCount, bitSums)
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).
			Message("Failed to scan.")
		os.Exit(1)
	}

	log.Info().WithStringf("sums", "%v", bitSums).
		WithInt("scans", scansCount).
		Message("Scanning complete.")

	var gamma int
	halfScansCount := scansCount / 2
	for i := 0; i < len(bitSums); i++ {
		gamma <<= 1
		if bitSums[i] > halfScansCount {
			gamma++
		}
		log.Debug().WithInt("idx", i).
			WithInt("sum", bitSums[i]).
			WithStringf("gamma", "%b", gamma).
			Message("")
	}

	epsilon := ((1 << len(bitSums)) - 1) & (^gamma)

	log.Info().
		WithInt("gamma", gamma).
		WithInt("epsilon", epsilon).
		WithInt("power", gamma*epsilon).
		Message("Calculated consumption.")
}

func addBits(bits int, bitCount int, sums []int) {
	for i := 0; i < bitCount; i++ {
		sums[bitCount-i-1] += (bits >> i) & 1
	}
}

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day03")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	scanner := NewBitScanner(inputFile)
	var bitSums []int
	var scansCount int
	for scanner.Scan() {
		bits := scanner.Bits()
		bitCount := scanner.BitCount()
		scansCount++
		bitSums = growSliceToLen(bitSums, bitCount)

		for i := 0; i < bitCount; i++ {
			bitSums[bitCount-i-1] += (bits >> i) & 1
		}
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

func growSliceToLen(slice []int, min int) []int {
	for len(slice) < min {
		slice = append(slice, 0)
	}
	return slice
}

func parseBits(s string) (int, int, error) {
	i, err := strconv.ParseInt(s, 2, 0)
	return int(i), len(s), err
}

func NewBitScanner(reader io.Reader) *bitScanner {
	return &bitScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type bitScanner struct {
	scanner  *bufio.Scanner
	num      int
	bitCount int
	err      error
}

func (bs *bitScanner) Scan() bool {
	if !bs.scanner.Scan() {
		return false
	}
	str := bs.scanner.Text()
	bs.num, bs.bitCount, bs.err = parseBits(str)
	log.Debug().
		WithString("str", str).
		WithInt("count", bs.bitCount).
		WithStringf("bits", "0b%b", bs.num).
		Messagef("%T", bs)
	return bs.err == nil
}

func (bs *bitScanner) Bits() int {
	return bs.num
}

func (bs *bitScanner) BitCount() int {
	return bs.bitCount
}

func (bs *bitScanner) Err() error {
	return bs.err
}

package main

import (
	"bufio"
	"io"
	"strconv"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day03")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	if common.Part2 {
		part2(inputFile)
	} else {
		part1(inputFile)
	}
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

func (bs *bitScanner) BitsCount() int {
	return bs.bitCount
}

func (bs *bitScanner) Err() error {
	return bs.err
}

package main

import (
	"os"
	"strconv"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day01")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	var scanner = NewIntScanner(inputFile)
	if common.Part2 {
		scanner = NewWindowedScanner(scanner, 3)
		log.Info().WithInt("window", 3).Message("Using windowed scanner.")
	}
	var (
		prevDepth int
		increases int
		scans     int
	)

	for scanner.Scan() {
		depth := scanner.Int()
		scans++
		if prevDepth != 0 && depth > prevDepth {
			increases++
		}
		prevDepth = depth
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).
			Message("Failed to scan.")
		os.Exit(1)
	}

	log.Info().WithInt("scans", scans).WithInt("increases", increases).
		Message("Scanning complete.")
}

type IntScanner interface {
	Scan() bool
	Int() int
	Err() error
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

package main

import (
	"os"
	"strconv"

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

	var uniqueSegDigits int
	for _, entry := range entries {
		for _, out := range entry.OutputValue {
			switch len(out) {
			case 2, 3, 4, 7:
				uniqueSegDigits++
			}
		}
	}
	log.Info().WithInt("digits", uniqueSegDigits).
		Message("Counted unique segment-count digits in outputs.")
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

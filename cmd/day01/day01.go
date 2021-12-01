package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/logger/consolepretty"
	"github.com/spf13/pflag"
)

var log = logger.New()

func main() {
	pflag.Usage = func() {
		fmt.Printf(`Advent of Code 2021 solution: '%s'

Flags:
`, filepath.Base(os.Args[0]))
		pflag.PrintDefaults()
	}
	var inputPath string
	var part2 bool
	var showHelp bool
	var showDebug bool
	pflag.StringVarP(&inputPath, "input", "i", "input.txt", "Puzzle input")
	pflag.BoolVarP(&part2, "part2", "2", false, "Give part 2 results")
	pflag.BoolVarP(&showDebug, "verbose", "v", false, "Show debug text")
	pflag.BoolVarP(&showHelp, "help", "h", false, "Show this help text")
	pflag.Parse()

	prettyConfig := consolepretty.DefaultConfig
	prettyConfig.ScopeMinLengthAuto = false
	prettyConfig.CallerMinLength = 15
	prettyConfig.CallerMaxLength = 15
	prettyConfig.DisableDate = true
	prettyConsole := consolepretty.New(prettyConfig)
	if showDebug {
		logger.AddOutput(logger.LevelDebug, prettyConsole)
	} else {
		logger.AddOutput(logger.LevelInfo, prettyConsole)
	}

	if showHelp {
		log.Debug().Message("Help requested.")
		pflag.Usage()
		os.Exit(0)
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		log.Error().WithError(err).WithString("path", inputPath).
			Message("Failed to open input file.")
		os.Exit(1)
	} else {
		log.Info().WithString("path", inputPath).Message("Reading file.")
	}
	defer inputFile.Close()

	var scanner = NewIntScanner(inputFile)
	if part2 {
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

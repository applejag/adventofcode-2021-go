package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/logger/consolepretty"
	"github.com/spf13/pflag"
)

var (
	InputPath string
	Part2     bool
	ShowHelp  bool
	ShowDebug bool
)

var log = logger.NewScoped("common")

func init() {
	pflag.StringVarP(&InputPath, "input", "i", "input.txt", "Puzzle input")
	pflag.BoolVarP(&Part2, "part2", "2", false, "Give part 2 results")
	pflag.BoolVarP(&ShowDebug, "verbose", "v", false, "Show debug text")
	pflag.BoolVarP(&ShowDebug, "help", "h", false, "Show this help text")

	pflag.Usage = func() {
		fmt.Printf(`Advent of Code 2021 solution: '%s'

Flags:
`, filepath.Base(os.Args[0]))
		pflag.PrintDefaults()
	}
}

func Init() {
	pflag.Parse()
	initLogger()

	if ShowHelp {
		log.Debug().Message("Help requested.")
		pflag.Usage()
		os.Exit(0)
	}
}

func initLogger() {
	if !pflag.Parsed() {
		panic("Must parse pflags first!")
	}

	prettyConfig := consolepretty.DefaultConfig
	prettyConfig.CallerMinLength = 15
	prettyConfig.CallerMaxLength = 15
	prettyConfig.DisableDate = true
	prettyConsole := consolepretty.New(prettyConfig)
	if ShowDebug {
		logger.AddOutput(logger.LevelDebug, prettyConsole)
	} else {
		logger.AddOutput(logger.LevelInfo, prettyConsole)
	}
}

func OpenInput() *os.File {
	inputFile, err := os.Open(InputPath)
	if err != nil {
		log.Error().WithError(err).WithString("path", InputPath).
			Message("Failed to open input file.")
		os.Exit(1)
	} else {
		log.Info().WithString("path", InputPath).Message("Reading file.")
	}
	return inputFile
}

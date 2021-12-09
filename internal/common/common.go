package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/logger/consolejson"
	"github.com/iver-wharf/wharf-core/pkg/logger/consolepretty"
	"github.com/spf13/pflag"
)

var (
	InputPath    string
	OutputFormat string
	Part2        bool
	ShowHelp     bool
	ShowDebug    bool
)

var log = logger.NewScoped("common")

func init() {
	pflag.StringVarP(&OutputFormat, "output", "o", "pretty", "Output format, 'pretty' or 'json'")
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

	part := 1
	if Part2 {
		part = 2
	}
	log.Debug().WithInt("part", part).Message("Computing:")
}

func initLogger() {
	if !pflag.Parsed() {
		panic("Must parse pflags first!")
	}

	switch strings.ToLower(OutputFormat) {
	case "pretty":
		prettyConfig := consolepretty.DefaultConfig
		prettyConfig.CallerMinLength = 15
		prettyConfig.CallerMaxLength = 15
		prettyConfig.DisableDate = true
		prettyOut := consolepretty.New(prettyConfig)
		if ShowDebug {
			logger.AddOutput(logger.LevelDebug, prettyOut)
		} else {
			logger.AddOutput(logger.LevelInfo, prettyOut)
		}
	case "json":
		jsonConfig := consolejson.Config{
			TimeFormat: consolejson.TimeRFC3339,
		}
		jsonOut := consolejson.New(jsonConfig)
		if ShowDebug {
			logger.AddOutput(logger.LevelDebug, jsonOut)
		} else {
			logger.AddOutput(logger.LevelInfo, jsonOut)
		}
	default:
		fmt.Println("Invalid output format:", OutputFormat)
		pflag.Usage()
		os.Exit(1)
	}
}

func OpenInput() *os.File {
	abspath, err := filepath.Abs(InputPath)
	if err != nil {
		log.Error().WithError(err).WithString("path", InputPath).
			Message("Failed to resolve input path name.")
		os.Exit(1)
	}
	log.Debug().WithString("abspath", abspath).Message("")

	inputFile, err := os.Open(abspath)
	if err != nil {
		log.Error().WithError(err).WithString("path", InputPath).
			Message("Failed to open input file.")
		os.Exit(1)
	}

	log.Info().WithString("path", InputPath).Message("Reading file.")
	return inputFile
}

package main

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day06")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	bytes, err := io.ReadAll(inputFile)
	if err != nil {
		log.Error().WithError(err).Message("Failed to read file.")
		os.Exit(1)
	}

	numsStrs := strings.Split(strings.TrimSpace(string(bytes)), ",")
	fishAges, err := parseInts(numsStrs)
	if err != nil {
		log.Error().WithError(err).Message("Failed to parse int.")
		os.Exit(1)
	}

	log.Info().WithInt("fish", len(fishAges)).
		Message("Scanning complete.")

	for day := 0; day < 80; day++ {
		fishAges = iterateAges(fishAges)
	}
	log.Info().WithInt("fish", len(fishAges)).
		Message("Multiplied.")
}

func iterateAges(fishAges []int) []int {
	for i, age := range fishAges {
		if age <= 0 {
			fishAges = append(fishAges, 8)
			age = 7
		}
		fishAges[i] = age - 1
	}
	return fishAges
}

func parseInts(s []string) ([]int, error) {
	var ints []int
	for _, str := range s {
		i, err := parseInt(str)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

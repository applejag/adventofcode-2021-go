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
	nums, err := parseInts(numsStrs)
	if err != nil {
		log.Error().WithError(err).Message("Failed to parse int.")
		os.Exit(1)
	}

	log.Info().WithInt("fish", len(nums)).
		Message("Scanning complete.")

	var fishCountByAge [9]int64
	for _, age := range nums {
		fishCountByAge[age]++
	}

	limit := 80
	if common.Part2 {
		limit = 256
	}

	for day := 0; day < limit; day++ {
		log.Debug().
			WithInt64("0", fishCountByAge[0]).
			WithInt64("1", fishCountByAge[1]).
			WithInt64("2", fishCountByAge[2]).
			WithInt64("3", fishCountByAge[3]).
			WithInt64("4", fishCountByAge[4]).
			WithInt64("5", fishCountByAge[5]).
			WithInt64("6", fishCountByAge[6]).
			WithInt64("7", fishCountByAge[7]).
			WithInt64("8", fishCountByAge[8]).
			Message("Fish count by age.")
		iterateAges(&fishCountByAge)
	}

	var sum int64
	for _, fishCount := range fishCountByAge {
		sum += fishCount
	}

	log.Info().WithInt64("fish", sum).
		Message("Multiplied.")
}

func iterateAges(fishCountByAge *[9]int64) {
	ages0 := fishCountByAge[0]
	fishCountByAge[0] = fishCountByAge[1]
	for i := 0; i <= 7; i++ {
		fishCountByAge[i] = fishCountByAge[i+1]
	}
	fishCountByAge[6] += ages0
	fishCountByAge[8] = ages0
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

package main

import (
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day07")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	inputBytes, err := io.ReadAll(inputFile)
	if err != nil {
		log.Error().WithError(err).Message("Failed to read file content.")
		os.Exit(1)
	}

	positions, err := parseInts(strings.Split(strings.TrimSpace(string(inputBytes)), ","))
	if err != nil {
		log.Error().WithError(err).Message("Failed to parse int.")
		os.Exit(1)
	}

	log.Info().WithInt("crabs", len(positions)).
		Message("Scanning complete.")

	target := median(positions)
	log.Debug().
		WithInt("median", median(positions)).
		Message("")

	var fuelSum int
	for _, pos := range positions {
		delta := target - pos
		if delta < 0 {
			delta = -delta
		}
		fuelSum += delta
	}

	log.Info().WithInt("fuel", fuelSum).
		Message("Calculated fuel consumption.")
}

func median(nums []int) int {
	size := len(nums)
	halfSize := size / 2
	clone := make([]int, size)
	copy(clone, nums)
	sort.Ints(clone)

	if size%2 == 1 {
		return clone[halfSize]
	}

	return (clone[halfSize-1] + clone[halfSize]) / 2
}

func sumInts(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
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

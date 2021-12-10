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

var log = logger.NewScoped("day09")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	inputBytes, err := io.ReadAll(inputFile)
	if err != nil {
		log.Error().WithError(err).Message("Failed to read file.")
		os.Exit(1)
	}
	inputLines := strings.Split(strings.TrimSpace(string(inputBytes)), "\n")

	h, err := ParseHeightmap(inputLines)
	if err != nil {
		log.Error().WithError(err).Message("Failed to parse heightmap.")
		os.Exit(1)
	}

	log.Info().WithStringf("size", "%dx%d", len(h), len(h[0])).
		Message("Scanning complete.")

	if common.Part2 {
		sizes := h.GetBasinSizes()
		sort.Ints(sizes)
		larger3 := sizes[len(sizes)-3:]
		log.Debug().WithStringf("sizes", "%v", larger3).Message("Taking 3 largest.")
		product := 1
		for _, s := range larger3 {
			product *= s
		}
		log.Info().WithInt("product", product).Message("Multiplied basin sizes.")
	} else {
		sum := h.SumRiskLevels()
		log.Info().WithInt("sum", sum).Message("Summed lowest points.")
	}
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

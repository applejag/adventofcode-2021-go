package main

import (
	"os"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day11")

func main() {
	common.Init()

	inputLines := common.ReadInputLines()
	g, err := ParseGrid(inputLines)
	if err != nil {
		log.Error().WithError(err).Message("Failed to parse grid.")
		os.Exit(1)
	}

	w, h := g.Size()
	squidCount := w * h
	log.Info().WithStringf("size", "%dx%d", w, h).
		Message("Scanning complete.")

	log.Debug().Messagef("Grid:\n%s", g)

	totalSteps := 100
	if common.Part2 {
		// better than keeping on forever
		totalSteps = 10000
	}

	var flashesSum int
	for step := 1; step <= totalSteps; step++ {
		f := g.Iterate()

		if common.Part2 && f == squidCount {
			log.Info().WithInt("step", step).Message("All squids flashed!")
			return
		}

		flashesSum += f
		log.Debug().WithInt("step", step).
			WithInt("flashes", f).
			WithInt("sum", flashesSum).
			Message("")
		if step <= 10 || step%10 == 0 || step == totalSteps {
			log.Debug().Messagef("Grid:\n%s", g)
		}
	}

	if common.Part2 {
		log.Warn().WithInt("steps", totalSteps).
			Message("All squids never fired at the same time.")
	} else {
		log.Info().WithInt("steps", totalSteps).WithInt("sum", flashesSum).
			Message("Counted flashes from 100 steps.")
	}
}

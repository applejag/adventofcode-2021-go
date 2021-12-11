package main

import (
	"os"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
	"github.com/spf13/pflag"
)

var log = logger.NewScoped("day11")

func main() {
	iterations := 100
	pflag.IntVarP(&iterations, "steps", "s", iterations, "number of iteration steps")
	common.Init()

	inputLines := common.ReadInputLines()
	g, err := ParseGrid(inputLines)
	if err != nil {
		log.Error().WithError(err).Message("Failed to parse grid.")
		os.Exit(1)
	}

	w, h := g.Size()
	log.Info().WithStringf("size", "%dx%d", w, h).
		Message("Scanning complete.")

	log.Debug().Messagef("Grid:\n%s", g)

	var flashesSum int
	for step := 0; step < iterations; step++ {
		f := g.Iterate()
		flashesSum += f
		log.Debug().WithInt("step", step).
			WithInt("flashes", f).
			WithInt("sum", flashesSum).
			Message("")
		if step < 10 || step%10 == 9 || step == iterations-1 {
			log.Debug().Messagef("Grid:\n%s", g)
		}
	}

	log.Info().WithInt("sum", flashesSum).Message("Counted flashes from 100 steps.")
}

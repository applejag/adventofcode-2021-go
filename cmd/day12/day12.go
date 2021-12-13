package main

import (
	"os"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day12")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	var scanner = newPathScanner(inputFile)
	var scans int
	m := Map{}

	for scanner.Scan() {
		path := scanner.Path()
		// set up bidirectional mappings
		m[path[0]] = append(m[path[0]], path[1])
		m[path[1]] = append(m[path[1]], path[0])
		scans++
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).
			Message("Failed to scan.")
		os.Exit(1)
	}

	log.Debug().
		WithStringf("start", "%v", m[CaveStart]).
		WithStringf("end", "%v", m[CaveEnd]).
		Message("Connections.")

	log.Info().WithInt("scans", scans).
		Message("Scanning complete.")

	paths := m.CountPaths()
	log.Info().WithInt("paths", paths).
		Message("Counted possible paths.")
}

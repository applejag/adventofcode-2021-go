package main

import (
	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day02")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	scanner := NewCmdScanner(inputFile)
	var sub submarine
	var cmdCount int
	for scanner.Scan() {
		cmd := scanner.Command()
		arg := scanner.Argument()
		log.Debug().WithString("cmd", string(cmd)).WithInt("arg", arg).Message("Scan.")

		sub.Move(cmd, arg)
		cmdCount++
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).Message("Failed to scan.")
	}
	log.Info().WithInt("cmds", cmdCount).Message("Scanning complete.")
	log.Info().
		WithInt("pos", sub.position).
		WithInt("depth", sub.depth).
		WithInt("product", sub.position*sub.depth).
		Message("Final submarine position.")
}

type Command string

const (
	CommandForward Command = "forward"
	CommandDown    Command = "down"
	CommandUp      Command = "up"
)

type submarine struct {
	depth    int
	position int
}

func (s *submarine) Move(cmd Command, arg int) {
	switch cmd {
	case CommandForward:
		s.position += arg
	case CommandDown:
		s.depth += arg
	case CommandUp:
		s.depth -= arg
	default:
		log.Warn().WithString("cmd", string(cmd)).Message("Unknown command.")
	}
}

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
	var sub Submarine
	if common.Part2 {
		sub = &submarine2{}
	} else {
		sub = &submarine1{}
	}
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
		WithInt("pos", sub.Position()).
		WithInt("depth", sub.Depth()).
		WithInt("product", sub.Position()*sub.Depth()).
		Message("Final submarine position.")
}

type Command string

const (
	CommandForward Command = "forward"
	CommandDown    Command = "down"
	CommandUp      Command = "up"
)

type Submarine interface {
	Move(Command, int)
	Depth() int
	Position() int
}

type submarine1 struct {
	depth    int
	position int
}

func (s *submarine1) Move(cmd Command, arg int) {
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

func (s *submarine1) Depth() int {
	return s.depth
}

func (s *submarine1) Position() int {
	return s.position
}

type submarine2 struct {
	depth    int
	position int
	aim      int
}

func (s *submarine2) Move(cmd Command, arg int) {
	switch cmd {
	case CommandForward:
		s.position += arg
		s.depth += s.aim * arg
	case CommandDown:
		s.aim += arg
	case CommandUp:
		s.aim -= arg
	default:
		log.Warn().WithString("cmd", string(cmd)).Message("Unknown command.")
	}
}

func (s *submarine2) Depth() int {
	return s.depth
}

func (s *submarine2) Position() int {
	return s.position
}

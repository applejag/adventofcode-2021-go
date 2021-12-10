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

var log = logger.NewScoped("day10")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	inputBytes, err := io.ReadAll(inputFile)
	if err != nil {
		log.Error().WithError(err).Message("Failed to read file.")
		os.Exit(1)
	}

	inputLines := strings.Split(string(inputBytes), "\n")

	log.Info().WithInt("lines", len(inputLines)).
		Message("Scanning complete.")

	var syntaxScore int
	var corruptChunks int

	var autocompleteScores []int

lineLoop:
	for lineIdx, line := range inputLines {
		var stack ChunkRuneStack
		for rIdx, r := range line {
			logFunc := func(e logger.Event) logger.Event {
				return e.WithStringf("pos", "%d:%d", lineIdx+1, rIdx).WithString("rune", string(r))
			}
			cr, ok := ParseChunkRune(r)
			if !ok {
				log.Error().WithFunc(logFunc).Message("Failed to parse chunk rune.")
				os.Exit(1)
			}
			if cr.open {
				stack.Push(cr)
				continue
			}
			top, ok := stack.Pop()
			if !ok {
				log.Error().WithFunc(logFunc).Message("Closing chunk but stack is empty.")
				os.Exit(1)
			}
			if top.Chunk == cr.Chunk {
				continue
			}
			if common.Part2 {
				continue lineLoop
			}
			score := cr.Chunk.SyntaxErrScore()
			log.Debug().WithFunc(logFunc).
				WithString("expected", top.Chunk.Close()).
				WithStringf("score", "%+d", score).
				Message("Chunk is corrupt.")
			syntaxScore += score
			corruptChunks++
			break
		}

		if !common.Part2 {
			continue
		}
		if len(stack) == 0 {
			continue
		}
		var score int
		for len(stack) > 0 {
			cr, _ := stack.Pop()
			score = score*5 + cr.Chunk.AutocompleteScore()
		}
		autocompleteScores = append(autocompleteScores, score)
		log.Debug().
			WithStringf("pos", "%d:%d", lineIdx+1, len(line)).
			WithInt("score", score).
			Message("Chunk is incomplete.")
	}

	if common.Part2 {
		sort.Ints(autocompleteScores)
		median := autocompleteScores[len(autocompleteScores)/2]
		log.Info().
			WithInt("chunks", len(autocompleteScores)).
			WithInt("score", median).
			Message("Found median incomplete chunk score.")
	} else {
		log.Info().
			WithInt("chunks", corruptChunks).
			WithInt("score", syntaxScore).
			Message("Summed corrupting chunk rune scores.")
	}
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

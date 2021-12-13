package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day13")

func main() {
	common.Init()

	inputLines := common.ReadInputLines()

	var foldLines []string
	var points []Point
	for i, line := range inputLines {
		if line == "" {
			foldLines = inputLines[i+1:]
			break
		}
		var p Point
		_, err := fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		if err != nil {
			log.Error().WithError(err).Message("Failed to scan point.")
			os.Exit(1)
		}
		points = append(points, p)
	}

	var folds []Fold
	for _, line := range foldLines {
		var f Fold
		_, err := fmt.Sscanf(line, "fold along %c=%d", &f.Axis, &f.Pos)
		if err != nil {
			log.Error().WithError(err).Message("Failed to scan fold.")
			os.Exit(1)
		}
		folds = append(folds, f)
	}

	log.Info().
		WithInt("points", len(points)).
		WithInt("folds", len(folds)).
		Message("Scanning complete.")

	if common.Part2 {
		for _, fold := range folds {
			points = foldPaper(points, fold)
		}
		log.Info().
			Message(paperString(points))
	} else {
		folded := foldPaper(points, folds[0])
		log.Info().
			WithInt("points", len(folded)).
			Message("Folded once.")
	}
}

func paperString(points []Point) string {
	var maxX, maxY int
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	var sb strings.Builder
	sb.WriteString("Paper:")
	for y := 0; y <= maxY; y++ {
		sb.WriteByte('\n')
		for x := 0; x <= maxX; x++ {
			if containsPoint(points, Point{x, y}) {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	return sb.String()
}

func foldPaper(points []Point, fold Fold) []Point {
	var result []Point
	for _, p := range points {
		p2 := p.Folded(fold)
		if !containsPoint(result, p2) {
			result = append(result, p2)
		}
	}
	return result
}

func containsPoint(slice []Point, elem Point) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

type Fold struct {
	Axis Axis
	Pos  int
}

type Axis rune

const (
	AxisX Axis = 'x'
	AxisY Axis = 'y'
)

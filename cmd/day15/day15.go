package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day15")

func main() {
	common.Init()

	inputLines := common.ReadInputLines()

	g, err := ParseGrid(inputLines)
	if err != nil {
		log.Error().WithError(err).
			Message("Failed to parse grid.")
		os.Exit(1)
	}

	w, h := g.Size()
	log.Info().WithStringf("size", "%dx%d", w, h).
		Message("Scanning complete.")

	scored := createScoredGrid(g)
	//debugPrintGrid(scored)
	path := findPath(scored, Point{0, 0})
	debugPrintPath(g, path)
	risk := scored.Val(path.Points[0])
	log.Info().WithInt("risk", risk).Message("Found optimal path")
}

func debugPrintGrid(g Grid) {
	if !common.ShowDebug {
		return
	}
	max := g.MaxVal()
	intWidth := int(math.Log10(float64(max))) + 1
	format := fmt.Sprintf("%%%dd ", intWidth)
	w, h := g.Size()
	var sb strings.Builder
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := Point{x, y}
			fmt.Fprintf(&sb, format, g.Val(p))
		}
		sb.WriteByte('\n')
	}
	fmt.Print(sb.String())
}

func debugPrintPath(g Grid, path Path) {
	if !common.ShowDebug {
		return
	}
	var sb strings.Builder
	w, h := g.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := Point{x, y}
			if path.ContainsPoint(p) {
				sb.WriteRune(g.Rune(p))
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Print(sb.String())
}

func createScoredGrid(g Grid) Grid {
	w, h := g.Size()
	scored := NewGrid(w, h)
	for x := w - 1; x >= 0; x-- {
		for y := h - 1; y >= 0; y-- {
			p := Point{x, y}
			neighbors := scored.GetNeighbors2SouthEast(p)
			if len(neighbors) == 0 {
				scored.Set(p, g.Val(p))
			} else {
				neighborVals := scored.Vals(neighbors)
				sort.Ints(neighborVals)
				scored.Set(p, g.Val(p)+neighborVals[0])
			}
		}
	}
	return scored
}

func findPath(g Grid, p Point) Path {
	return findPathRec(g, p, Path{})
}

type Path struct {
	Points []Point
	Sum    int
}

func (path Path) ContainsPoint(point Point) bool {
	for _, p := range path.Points {
		if p == point {
			return true
		}
	}
	return false
}

func findPathRec(g Grid, p Point, path Path) Path {
	neighbors := g.GetNeighbors2SouthEast(p)
	if len(neighbors) == 0 {
		return path
	}
	neighborVals := g.Vals(neighbors)
	sort.Ints(neighborVals)
	//log.Debug().
	//	WithStringer("point", p).
	//	WithStringf("neighbors", "%v", neighborVals).
	//	Message("")
	lowestVal := neighborVals[0]
	var paths []Path
	for _, n := range neighbors {
		val := g.Val(n)
		if val == lowestVal {
			paths = append(paths, findPathRec(g, n, Path{append(path.Points, n), path.Sum + val}))
		}
	}
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Sum < paths[j].Sum
	})
	return paths[0]
}

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
	"github.com/jilleJr/adventofcode-2021-go/internal/util"
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

	path, err := findPathDijkstra(g, Point{w - 1, h - 1}, Point{0, 0})
	if err != nil {
		log.Error().WithError(err).Message("Failed to calculate path.")
		os.Exit(1)
	}
	debugPrintPath(g, path)
	log.Debug().WithStringf("path", "%v", path).Message("")
	risk := util.SumInts(g.Vals(path.Points))
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

func findPathDijkstra(g Grid, source Point, target Point) (Path, error) {
	nodes := map[Point]int{}
	w, h := g.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			p := Point{x, y}
			nodes[p] = math.MaxInt
		}
	}
	nodes[source] = 0
	prevPoints := map[Point]Point{}

	for len(nodes) > 0 {
		minPoint := Point{}
		minNodeDist := math.MaxInt
		for p, nDist := range nodes {
			if nDist < minNodeDist {
				minNodeDist = nDist
				minPoint = p
			}
		}

		delete(nodes, minPoint)

		neighbors := g.GetNeighbors4(minPoint)
		//log.Debug().
		//	WithInt("remaining", len(nodes)).
		//	WithStringf("node", "%v", minPoint).
		//	WithStringf("neighbors", "%v", neighbors).
		//	Message("Dijkstra pathing.")
		for _, neighborPoint := range neighbors {
			if _, hasPoint := nodes[neighborPoint]; !hasPoint {
				continue
			}
			dist := minNodeDist + g.Val(neighborPoint)
			if dist < nodes[neighborPoint] {
				nodes[neighborPoint] = dist
				prevPoints[neighborPoint] = minPoint
			}
		}
	}
	p, hasPoint := prevPoints[target]
	if !hasPoint && target != source {
		return Path{}, errors.New("target is unreachable")
	}
	var path []Point
	for hasPoint {
		path = append(path, p)
		p, hasPoint = prevPoints[p]
	}
	score := util.SumInts(g.Vals(path))
	return Path{path, score}, nil
}

package main

import (
	"fmt"
	"strings"
)

type Empty struct{}

func parseDigit(d rune) (int, error) {
	if d < '0' || d > '9' {
		return 0, fmt.Errorf("invalid digit: %s", string(d))
	}
	return int(d - '0'), nil
}

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func NewGrid(w, h int) Grid {
	lines := make([][]int, 0, w)
	for x := 0; x < w; x++ {
		lines = append(lines, make([]int, h))
	}
	return Grid(lines)
}

func ParseGrid(strs []string) (Grid, error) {
	var lines [][]int
	for _, str := range strs {
		var line []int
		for _, r := range str {
			d, err := parseDigit(r)
			if err != nil {
				return Grid{}, err
			}
			line = append(line, d)
		}
		lines = append(lines, line)
	}
	return Grid(lines), nil
}

type Grid [][]int

func (g Grid) String() string {
	var sb strings.Builder
	w, h := g.Size()
	for x := 0; x < w; x++ {
		if x > 0 {
			sb.WriteByte('\n')
		}
		for y := 0; y < h; y++ {
			sb.WriteRune(g.Rune(Point{x, y}))
		}
	}
	return sb.String()
}

func (g Grid) Rune(p Point) rune {
	val := g.Val(p)
	if val > 9 {
		return '?'
	}
	return '0' + rune(val)
}

func (g Grid) Val(p Point) int {
	return g[p.y][p.x]
}

func (g Grid) Vals(points []Point) []int {
	vals := make([]int, len(points))
	for i, p := range points {
		vals[i] = g.Val(p)
	}
	return vals
}

func (g Grid) MaxVal() int {
	w, h := g.Size()
	var max int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			val := g.Val(Point{x, y})
			if val > max {
				max = val
			}
		}
	}
	return max
}

func (g Grid) Set(p Point, val int) {
	g[p.y][p.x] = val
}

func (g Grid) Inc(p Point) int {
	g[p.y][p.x]++
	return g[p.y][p.x]
}

func (g Grid) Size() (w int, h int) {
	if len(g) == 0 {
		return 0, 0
	}
	return len(g[0]), len(g)
}

func (g Grid) GetNeighbors2SouthEast(p Point) []Point {
	var neighbors []Point
	w, h := g.Size()
	if p.x < w-1 {
		neighbors = append(neighbors, Point{p.x + 1, p.y})
	}
	if p.y < h-1 {
		neighbors = append(neighbors, Point{p.x, p.y + 1})
	}
	return neighbors
}

func (g Grid) GetNeighbors4(p Point) []Point {
	var neighbors []Point
	w, h := g.Size()
	if p.x > 0 {
		neighbors = append(neighbors, Point{p.x - 1, p.y})
	}
	if p.x < w-1 {
		neighbors = append(neighbors, Point{p.x + 1, p.y})
	}
	if p.y > 0 {
		neighbors = append(neighbors, Point{p.x, p.y - 1})
	}
	if p.y < h-1 {
		neighbors = append(neighbors, Point{p.x, p.y + 1})
	}
	return neighbors
}

func (g Grid) GetNeighbors8(p Point) []Point {
	minX := p.x - 1
	maxX := p.x + 1
	minY := p.y - 1
	maxY := p.y + 1
	w, h := g.Size()
	if minX < 0 {
		minX = 0
	}
	if maxX >= w {
		maxX = w - 1
	}
	if minY < 0 {
		minY = 0
	}
	if maxY >= h {
		maxY = h - 1
	}
	var neighbors []Point
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if x == p.x && y == p.y {
				continue
			}
			neighbors = append(neighbors, Point{x, y})
		}
	}
	return neighbors
}

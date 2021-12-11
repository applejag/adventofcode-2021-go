package main

import "strings"

type Point struct {
	x int
	y int
}

type Grid [][]byte

func (g Grid) String() string {
	var sb strings.Builder
	w, h := g.Size()
	for x := 0; x < w; x++ {
		if x > 0 {
			sb.WriteByte('\n')
		}
		for y := 0; y < h; y++ {
			sb.WriteByte('0' + g.Val(Point{x, y}))
		}
	}
	return sb.String()
}

func (g Grid) Val(p Point) byte {
	return g[p.x][p.y]
}

func (g Grid) Set(p Point, val byte) {
	g[p.x][p.y] = val
}

func (g Grid) Inc(p Point) byte {
	g[p.x][p.y]++
	return g[p.x][p.y]
}

func (g Grid) Size() (int, int) {
	return len(g), len(g[0])
}

func (g Grid) Iterate() int {
	alreadyFlashed := map[Point]Empty{}
	w, h := g.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			g.iteratePointRec(Point{x, y}, alreadyFlashed)
		}
	}
	return len(alreadyFlashed)
}

func (g Grid) iteratePointRec(p Point, alreadyFlashed map[Point]Empty) {
	if _, ok := alreadyFlashed[p]; ok {
		return
	}
	val := g.Inc(p)
	if val <= 9 {
		return
	}
	//log.Debug().WithStringf("pos", "%d,%d", p.x, p.y).Message("Flash!")
	alreadyFlashed[p] = Empty{}
	g.Set(p, 0)
	for _, n := range g.GetNeighbors(p) {
		g.iteratePointRec(n, alreadyFlashed)
	}
}

func (g Grid) GetNeighbors(p Point) []Point {
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

func ParseGrid(strs []string) (Grid, error) {
	var lines [][]byte
	for _, str := range strs {
		var line []byte
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

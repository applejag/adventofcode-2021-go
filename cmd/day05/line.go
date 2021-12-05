package main

import "fmt"

type Line struct {
	x1, y1, x2, y2 int
}

func (line Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", line.x1, line.y1, line.x2, line.y2)
}

func (line Line) Blit1(grid [][]int) {
	if line.x1 != line.x2 && line.y1 != line.y2 {
		// don't deal with diagonals
		return
	}
	if line.x2 < line.x1 {
		line.x2, line.x1 = line.x1, line.x2
	}
	if line.y2 < line.y1 {
		line.y2, line.y1 = line.y1, line.y2
	}

	log.Debug().
		WithStringer("line", line).
		WithStringf("grid", "%dx%d", len(grid), len(grid[0])).
		Message("Blit1")
	for x := line.x1; x <= line.x2; x++ {
		for y := line.y1; y <= line.y2; y++ {
			grid[x][y]++
		}
	}
}

func (line Line) Blit2(grid [][]int) {
	dx := line.x2 - line.x1
	dy := line.y2 - line.y1
	sdx := unit(dx)
	sdy := unit(dy)
	dist := max(abs(dx), abs(dy))

	log.Debug().
		WithStringer("line", line).
		WithInt("deltaX", dx).
		WithInt("deltaY", dy).
		WithStringf("grid", "%dx%d", len(grid), len(grid[0])).
		Message("Blit2")
	for d := 0; d <= dist; d++ {
		x := line.x1 + sdx*d
		y := line.y1 + sdy*d
		grid[x][y]++
	}
}

func unit(val int) int {
	switch {
	case val > 0:
		return 1
	case val < 0:
		return -1
	default:
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

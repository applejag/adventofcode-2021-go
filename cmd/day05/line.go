package main

import "fmt"

type Line struct {
	x1, y1, x2, y2 int
}

func (line Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", line.x1, line.y1, line.x2, line.y2)
}

func (line Line) Blit(grid [][]int) {
	if line.x1 != line.x2 && line.y1 != line.y2 {
		// don't deal with diagonals
		return
	}

	log.Debug().
		WithStringer("line", line).
		WithStringf("grid", "%dx%d", len(grid), len(grid[0])).
		Message("")
	for x := line.x1; x <= line.x2; x++ {
		for y := line.y1; y <= line.y2; y++ {
			grid[x][y]++
		}
	}
}

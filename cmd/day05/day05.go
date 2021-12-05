package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day05")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	scanner := NewLineScanner(inputFile)
	var lines []Line
	for scanner.Scan() {
		//line := scanner.Line()
		lines = append(lines, scanner.Line())
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).
			Message("Failed to scan.")
		os.Exit(1)
	}

	log.Info().WithInt("lines", len(lines)).
		Message("Scanning complete.")

	minMax := calcMinMax(lines)
	log.Debug().WithStringer("minmax", minMax).Message("Calculated line usage.")
	grid := make2D(minMax.x2+1, minMax.y2+1) //+1 because min/max are inclusive

	for _, line := range lines {
		line.Blit(grid)
	}

	if common.ShowDebug {
		var sb strings.Builder
		for y := minMax.y1; y <= minMax.y2; y++ {
			for x := minMax.x1; x <= minMax.x2; x++ {
				val := grid[x][y]
				if val == 0 {
					sb.WriteByte('.')
				} else {
					sb.WriteString(strconv.Itoa(val))
				}
			}
			sb.WriteByte('\n')
		}
		fmt.Println(sb.String())
	}

	var overlaps int
	for x := minMax.x1; x <= minMax.x2; x++ {
		for y := minMax.y1; y <= minMax.y2; y++ {
			if grid[x][y] >= 2 {
				overlaps++
			}
		}
	}

	log.Info().WithInt("overlaps", overlaps).Message("")
}

func make2D(width, height int) [][]int {
	slice := make([][]int, width)
	for x := 0; x < width; x++ {
		slice[x] = make([]int, height)
	}
	return slice
}

func calcMinMax(lines []Line) Line {
	minMax := lines[0]
	for _, line := range lines {
		if line.x1 < minMax.x1 {
			minMax.x1 = line.x1
		}
		if line.y1 < minMax.y1 {
			minMax.y1 = line.y1
		}
		if line.x2 > minMax.x2 {
			minMax.x2 = line.x2
		}
		if line.y2 > minMax.y2 {
			minMax.y2 = line.y2
		}
	}
	return minMax
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

func NewLineScanner(reader io.Reader) *lineScanner {
	return &lineScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type lineScanner struct {
	scanner *bufio.Scanner
	line    Line
	err     error
}

func (s *lineScanner) Scan() bool {
	if !s.scanner.Scan() {
		s.err = s.scanner.Err()
		return false
	}
	lineStr := s.scanner.Text()
	var x1, y1, x2, y2 int
	_, err := fmt.Sscanf(lineStr, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
	if err != nil {
		s.err = err
		return false
	}
	if x2 < x1 {
		x2, x1 = x1, x2
	}
	if y2 < y1 {
		y2, y1 = y1, y2
	}
	s.line = Line{x1, y1, x2, y2}
	log.Debug().
		WithInt("x1", x1).
		WithInt("y1", y1).
		WithInt("x2", x2).
		WithInt("y2", y2).
		Messagef("%T", s)
	return true
}

func (s *lineScanner) Line() Line {
	return s.line
}

func (s *lineScanner) Err() error {
	return s.err
}

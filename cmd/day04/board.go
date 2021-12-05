package main

import "strings"

type Cell struct {
	Num      int
	IsCalled bool
}

type Board struct {
	Cells [][]Cell
}

func ParseBoard(lines []string) (Board, error) {
	cells := make([][]Cell, 0, 5)
	for _, line := range lines {
		row := make([]Cell, 0, 5)
		for _, numStr := range strings.Split(line, " ") {
			if numStr == "" {
				continue
			}
			num, err := parseInt(numStr)
			if err != nil {
				return Board{}, err
			}
			row = append(row, Cell{
				Num: num,
			})
		}
		cells = append(cells, row)
	}
	return Board{
		Cells: cells,
	}, nil
}

func (b *Board) CallNumber(num int) bool {
	var anyCalled bool
	for _, row := range b.Cells {
		for i, cell := range row {
			if cell.Num == num {
				anyCalled = true
				row[i].IsCalled = true
			}
		}
	}
	return anyCalled
}

func (b *Board) HasWon() bool {
	for x := 0; x < 5; x++ {
		allCalled := true
		for y := 0; y < 5; y++ {
			if !b.Cells[x][y].IsCalled {
				allCalled = false
			}
		}
		if allCalled {
			return true
		}
	}
	for y := 0; y < 5; y++ {
		allCalled := true
		for x := 0; x < 5; x++ {
			if !b.Cells[x][y].IsCalled {
				allCalled = false
			}
		}
		if allCalled {
			return true
		}
	}
	return false
}

func (b *Board) SumUncalledNumbers() int {
	var sum int
	for _, row := range b.Cells {
		for _, cell := range row {
			if !cell.IsCalled {
				sum += cell.Num
			}
		}
	}
	return sum
}

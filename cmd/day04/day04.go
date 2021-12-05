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

var log = logger.NewScoped("day04")

func main() {
	common.Init()

	inputFile := common.OpenInput()
	defer inputFile.Close()

	scanner := newSectionScanner(inputFile)

	nums, err := scanNums(scanner)
	if err != nil {
		log.Error().WithError(err).
			Message("Failed to scan numbers.")
		os.Exit(1)
	}

	var bingoBoards []Board
	for scanner.Scan() {
		bingo, err := ParseBoard(scanner.Lines())
		if err != nil {
			log.Error().WithError(err).
				Message("Failed to parse bingo board.")
			os.Exit(1)
		}
		bingoBoards = append(bingoBoards, bingo)
	}
	if err := scanner.Err(); err != nil {
		log.Error().WithError(err).
			Message("Failed to scan boards.")
		os.Exit(1)
	}

	log.Info().WithInt("nums", len(nums)).
		WithInt("boards", len(bingoBoards)).
		Message("Scanning complete.")

	if common.Part2 {
		part2(bingoBoards, nums)
	} else {
		part1(bingoBoards, nums)
	}
}

func part1(bingoBoards []Board, nums []int) {
	winningBoardIdx, lastNumIdx := callNumbersUntilWinner(bingoBoards, []int{}, nums)
	if winningBoardIdx < 0 {
		log.Error().Message("No winning board after all numbers were called.")
		os.Exit(1)
	}

	sum := bingoBoards[winningBoardIdx].SumUncalledNumbers()
	log.Info().
		WithInt("num", nums[lastNumIdx]).
		WithInt("sum", sum).
		WithInt("score", sum*nums[lastNumIdx]).
		Message("Found winning board.")
}

func part2(bingoBoards []Board, nums []int) {
	var winningBoard *Board
	var numIdx int
	var blackList []int
	log.Debug().WithInt("boards", len(bingoBoards)).Message("Starting checks.")
	for numIdx < len(nums) {
		boardIdx, lastNumIdx := callNumbersUntilWinner(bingoBoards, blackList, nums[numIdx:])
		if boardIdx < 0 {
			break
		}
		log.Debug().WithInt("boardIdx", boardIdx).
			WithInt("numIdx", lastNumIdx).
			Message("Found another win.")
		numIdx = lastNumIdx + numIdx
		winningBoard = &bingoBoards[boardIdx]
		blackList = append(blackList, boardIdx)
	}
	if winningBoard == nil {
		log.Error().Message("No winning board after all numbers were called.")
		os.Exit(1)
	}

	sum := winningBoard.SumUncalledNumbers()
	log.Info().
		WithInt("num", nums[numIdx]).
		WithInt("sum", sum).
		WithInt("score", sum*nums[numIdx]).
		Message("Found last winning board.")
}

func callNumbersUntilWinner(boards []Board, boardBlackList, nums []int) (boardIdx, lastNumIdx int) {
	for numIdx, num := range nums {
		for boardIdx, board := range boards {
			if containsInt(boardBlackList, boardIdx) {
				continue
			}
			if board.CallNumber(num) && board.HasWon() {
				return boardIdx, numIdx
			}
		}
	}
	return -1, len(nums) - 1
}

func containsInt(slice []int, value int) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

func scanNums(scanner *sectionScanner) ([]int, error) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	numsStrs := scanner.Lines()
	if len(numsStrs) != 1 {
		return nil, fmt.Errorf("nums must only be 1 line, but was %d lines", len(numsStrs))
	}
	nums, err := parseNums(numsStrs[0])
	if err != nil {
		return nil, err
	}
	log.Debug().WithInt("count", len(nums)).
		WithStringf("nums", "%v", nums).
		Message("Parsed nums.")
	return nums, nil
}

func parseNums(s string) ([]int, error) {
	var nums []int
	var numStrs = strings.Split(s, ",")
	for _, numStr := range numStrs {
		num, err := parseInt(numStr)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

func newSectionScanner(reader io.Reader) *sectionScanner {
	return &sectionScanner{
		scanner: bufio.NewScanner(reader),
	}
}

type sectionScanner struct {
	scanner *bufio.Scanner
	lines   []string
	err     error
}

func (ss *sectionScanner) Scan() bool {
	ss.lines = nil
	for ss.scanner.Scan() {
		line := ss.scanner.Text()
		if line != "" {
			ss.lines = append(ss.lines, line)
		} else {
			return true
		}
	}
	ss.err = ss.scanner.Err()
	return ss.lines != nil && ss.err == nil
}

func (ss *sectionScanner) Lines() []string {
	return ss.lines
}

func (ss *sectionScanner) Err() error {
	return ss.err
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"
)

func main() {
	pflag.Usage = func() {
		fmt.Printf(`Advent of Code 2021 solution: '%s'

Flags:
`, filepath.Base(os.Args[0]))
		pflag.PrintDefaults()
	}
	var inputPath string
	var part2 bool
	var showHelp bool
	pflag.StringVarP(&inputPath, "input", "i", "input.txt", "Puzzle input")
	pflag.BoolVarP(&part2, "part2", "2", false, "Give part 2 results")
	pflag.BoolVarP(&showHelp, "help", "h", false, "Show this help text")
	pflag.Parse()

	if showHelp {
		pflag.Usage()
		os.Exit(0)
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var prevDepth int
	var increases int
	var scans int
	for scanner.Scan() {
		depth, err := parseInt(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		scans++
		if prevDepth != 0 && depth > prevDepth {
			increases++
		}
		prevDepth = depth
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Number of scans:", scans)
	log.Println("Number of depth increases:", increases)
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

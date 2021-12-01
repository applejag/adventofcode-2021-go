package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	var inputPath string
	flag.StringVar(&inputPath, "input", "input.txt", "Puzzle input")
	flag.Parse()

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

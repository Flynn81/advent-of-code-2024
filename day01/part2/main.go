package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "regexp"
	// "strconv"
)

//https://adventofcode.com/2023/day/1

func main() {
	readFile, err := os.Open("input")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	
	fmt.Println(sumSimilarityScores(lines))
}

func sumSimilarityScores(lines []string) int {
	total := 0
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))

	for i, line := range lines {
		left, right := processLine(line)
		leftList[i] = left
		rightList[i] = right
	}
	
	for i := 0; i < len(lines); i++ {
		total = total + similarityScore(leftList[i], rightList)
	}
	return total
}

func processLine(line string) (int, int) {
	parts := strings.Split(line, "   ")
	if parts == nil || len(parts) != 2 {
		panic("could not process line")
	}
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("could not convert first string to int")
	}
	right, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("could not convert second string to int")
	}
	return left, right
}

func similarityScore(left int, list[]int) int {
	score := 0
	for _,right := range list {
		if left == right {
			score = score + 1
		}
	}
	return score * left
}
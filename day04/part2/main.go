package main

import (
	"bufio"
	"fmt"
	"os"
	// "regexp"
	// "strconv"
)

var debug = false

func log(a ...any) {
	if debug {
		fmt.Println(a...)
	}
}

func main() {
	fmt.Println(readAndSearch("input"))
}

func readAndSearch(file string) int {
	readFile, err := os.Open(file)
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
	return count(lines)
}

func count(lines []string) int {
	total := 0

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			if string(lines[y][x]) == "A" {
				total = total + checkEveryDirection(lines, x, y)
			}
		}
	}
	return total
}

func checkEveryDirection(lines []string, x int, y int) int {
	if string(lines[y-1][x-1]) == "M" && string(lines[y+1][x+1]) == "S" {
		if (string(lines[y+1][x-1]) == "M" && string(lines[y-1][x+1]) == "S") ||
			(string(lines[y+1][x-1]) == "S" && string(lines[y-1][x+1]) == "M") {
			return 1
		}
	}
	if string(lines[y-1][x-1]) == "S" && string(lines[y+1][x+1]) == "M" {
		if (string(lines[y+1][x-1]) == "M" && string(lines[y-1][x+1]) == "S") ||
			(string(lines[y+1][x-1]) == "S" && string(lines[y-1][x+1]) == "M") {
			return 1
		}
	}
	return 0
}

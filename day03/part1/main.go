package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "regexp"
	// "strconv"
)

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

	fmt.Println(addUpAllResults(lines))
}

func addUpAllResults(lines []string) int {
	total := 0
	for _, line := range lines {
		total = total + processLine(line)
	}

	return total
}

func processLine(line string) int {
	total := 0
	var re = regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)
    
    for _, match := range re.FindAllString(line, -1) {
        total = total + multiply(match)
    }
	return total
}

func multiply(raw string) int {
	var re = regexp.MustCompile(`(?m)\d{1,3}`)
    
	left := 0
	right := 0
    for i, match := range re.FindAllString(raw, -1) {
        if i == 0 {
			left, _ = strconv.Atoi(match)
		} else if i == 1 {
			right, _ = strconv.Atoi(match)
		} else {
			panic("we shouldn't have more than 2 matches")
		}
    }
	return left * right
}

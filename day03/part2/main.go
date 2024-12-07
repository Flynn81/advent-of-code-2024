package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	instructions := ""
	for _, line := range lines {
		instructions = instructions + line
	}
	return processLine(instructions)
}

func processLine(line string) int {
	//firstDontIndex := strings.Index(line,"don't()")
	splits := strings.Split(line, "don't()")
	total := processSplit(splits[0])
	for i:=1; i<len(splits); i++ {
		if strings.Contains(splits[i], "do()") {
			doSplits := strings.Split(splits[i],"do()")
			for j := 1; j<len(doSplits); j++ {
				total = total + processSplit(doSplits[j])
			} 
		}
	}

	return total
}

func processSplit(split string) int {
	total := 0
	var re = regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)
    
    for _, match := range re.FindAllString(split, -1) {
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

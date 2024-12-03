package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	
	fmt.Println(sumDistances(lines)) //54953
}

func sumDistances(lines []string) int {
	total := 0
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))

	for i, line := range lines {
		left, right := processLine(line)
		leftList[i] = left
		rightList[i] = right
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	for i := 0; i < len(lines); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		total = total + distance
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

// func decodeLine(line string) int {
// 	// line = strings.ToUpper(line)
// 	// line = strings.ReplaceAll(line, "ONE", "O1E")
// 	// line = strings.ReplaceAll(line, "TWO", "T2O")
// 	// line = strings.ReplaceAll(line, "THREE", "TH3EE")
// 	// line = strings.ReplaceAll(line, "FOUR", "FO4R")
// 	// line = strings.ReplaceAll(line, "FIVE", "FI5E")
// 	// line = strings.ReplaceAll(line, "SIX", "S6X")
// 	// line = strings.ReplaceAll(line, "SEVEN", "SE7EN")
// 	// line = strings.ReplaceAll(line, "EIGHT", "EI8HT")
// 	// line = strings.ReplaceAll(line, "NINE", "N9NE")
// 	regex := regexp.MustCompile("[0-9]+")
// 	numbers := regex.FindAllString(line, -1)
// 	if numbers == nil {
// 		return 0
// 	}
// 	length := len(numbers)
// 	first := string(numbers[0][0])
// 	last := ""
// 	if length == 1 {
// 		if(len(numbers[0]) > 1) {
// 			last = string(numbers[0][len(numbers[length-1])-1])
// 		}
// 	} else {
// 		last = string(numbers[length-1][len(numbers[length-1])-1])
// 	}
// 	if last == "" {
// 		last = first
// 	}
// 	return convertStringToInt(first + last)
// }

// func convertStringToInt(s string) int {
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return i
// }
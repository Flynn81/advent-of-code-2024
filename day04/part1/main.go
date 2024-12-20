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

	// if check(lines, 0, -1, 0, 1, "XMAS") ||
	// 	check(lines, 0, -1, 0, 1, "SAMX") ||
	// 	check(lines, -1, -1, 1, 1, "SAMX") {
	// 	total = total + 1
	// }
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			total = total + checkEveryDirection(lines, "XMAS", x, y)
			total = total + checkEveryDirection(lines, "SAMX", x, y)
		}
	}
	return total
}

func checkEveryDirection(lines []string, word string, x int, y int) int {

	return check(lines, x, y, 1, 0, word) +
		check(lines, x, y, 0, 1, word) +
		check(lines, x, y, 1, 1, word) +
		check(lines, x, y, 1, -1, word)

}

func check(lines []string, x int, y int, vx int, vy int, word string) int {
	// if x == 4 && y == 4 && vy == 0 && vx == 1 {
	// 	debug = true
	// }
	log(string(lines[y][x]), string(word[0]), word)
	if lines[y][x] == word[0] {
		if len(word) == 1 {
			log("yoda1")
			debug = false
			return 1
		}
		log("yoda2")
		if x+vx >= len(lines[0]) || x+vx < 0 || y+vy >= len(lines) || y+vy < 0 {
			log("yoda 0")
			debug = false
			return 0
		}
		return check(lines, x+vx, y+vy, vx, vy, word[1:])
	}
	log("yoda3")
	debug = false
	return 0
}

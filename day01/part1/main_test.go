package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//processLine(line string) (int, int)
func TestProcessLine(t *testing.T) {
	left, right := processLine("1   2")
	assert.Equal(t, 1, left)
	assert.Equal(t, 2, right)

	left, right = processLine("13213213   465484646542")
	assert.Equal(t, 13213213, left)
	assert.Equal(t, 465484646542, right)
}

func TestSumDistances(t *testing.T) {
	readFile, err := os.Open("testInput")
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
	total := sumDistances(lines)
	assert.Equal(t, 11, total)
}
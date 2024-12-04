package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLineSafe(t *testing.T) {
	line := []int{7, 6, 4, 2, 1}
	isSafe := isLineSafe(line)
	assert.True(t, isSafe)

	line = []int{1, 2, 7, 8, 9}
	isSafe = isLineSafe(line)
	assert.False(t, isSafe)

	line = []int{9, 7, 6, 2, 1}
	isSafe = isLineSafe(line)
	assert.False(t, isSafe)

	line = []int{1, 3, 2, 4, 5}
	isSafe = isLineSafe(line)
	assert.False(t, isSafe)

	line = []int{8, 6, 4, 4, 1}
	isSafe = isLineSafe(line)
	assert.False(t, isSafe)

	line = []int{1, 3, 6, 7, 9}
	isSafe = isLineSafe(line)
	assert.True(t, isSafe)
}

func TestHowManySafe(t *testing.T) {
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
	total := howManySafe(lines)
	assert.Equal(t, 2, total)
}

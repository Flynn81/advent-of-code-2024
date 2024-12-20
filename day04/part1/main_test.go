package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	count := readAndSearch("testInput2")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput3")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput4")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput5")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput6")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput7")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput8")
	assert.Equal(t, 1, count)

	count = readAndSearch("testInput9")
	assert.Equal(t, 2, count)

	count = readAndSearch("testInput10")
	assert.Equal(t, 2, count)

	count = readAndSearch("testInput11")
	assert.Equal(t, 3, count)

	count = readAndSearch("testInput12")
	assert.Equal(t, 2, count)

	count = readAndSearch("testInput13")
	assert.Equal(t, 4, count)

	count = readAndSearch("testInput")
	assert.Equal(t, 18, count)
}

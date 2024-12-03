package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	result := multiply("mul(2,4)")
	assert.Equal(t, 8, result)
}

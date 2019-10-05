package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}

	for num%2 == 0 {
		num /= 2
	}

	for num%3 == 0 {
		num /= 3
	}

	for num%5 == 0 {
		num /= 5
	}

	return num == 1
}

func Test263(t *testing.T) {
	assert.Equal(t, false, isUgly(0))
	assert.Equal(t, true, isUgly(1))
	assert.Equal(t, true, isUgly(6))
	assert.Equal(t, true, isUgly(8))
	assert.Equal(t, false, isUgly(14))
}

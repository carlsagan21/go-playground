package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func missingNumber(nums []int) int {
	l := len(nums)
	total := (l*l + l) / 2
	for _, n := range nums {
		total -= n
	}
	return total
}

func Test268(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{3, 0, 1}))
	assert.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

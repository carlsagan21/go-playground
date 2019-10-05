package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxCount(m int, n int, ops [][]int) int {
	minA := m
	minB := n
	for _, op := range ops {
		if minA > op[0] {
			minA = op[0]
		}
		if minB > op[1] {
			minB = op[1]
		}
	}
	return minA * minB
}

func Test598(t *testing.T) {
	assert.Equal(t, 4, maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
}

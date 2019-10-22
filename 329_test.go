package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestIncreasingPath(matrix [][]int) int {
	longest := 0

	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			longest = max(longest, helper(i, j, matrix, cache))
		}
	}

	return longest
}

func helper(i int, j int, matrix [][]int, cache [][]int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return 0
	}

	if cache[i][j] != 0 {
		return cache[i][j]
	}

	longest := 0
	candidate := 0
	candidate = helper2(i-1, j, i, j, matrix, cache)
	longest = max(longest, candidate)
	candidate = helper2(i+1, j, i, j, matrix, cache)
	longest = max(longest, candidate)
	candidate = helper2(i, j-1, i, j, matrix, cache)
	longest = max(longest, candidate)
	candidate = helper2(i, j+1, i, j, matrix, cache)
	longest = max(longest, candidate)

	cache[i][j] = longest + 1
	return cache[i][j]
}

func helper2(ni int, nj int, i int, j int, matrix [][]int, cache [][]int) int {
	if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[0]) {
		return 0
	}

	if matrix[ni][nj] >= matrix[i][j] {
		return 0
	}

	return helper(ni, nj, matrix, cache)
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func Test329(t *testing.T) {
	assert.Equal(t, 3, longestIncreasingPath([][]int{{9, 6, 4}}))
	assert.Equal(t, 4, longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
}

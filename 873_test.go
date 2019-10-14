package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// I made index based DP, while value based DP is also possible.
func lenLongestFibSubseq(A []int) int {
	set := map[int]int{}
	for i, a := range A {
		set[a] = i
	}

	max := 0
	cache := make([][]int, len(A))
	for i := range cache {
		cache[i] = make([]int, len(A))
	}

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			hValue := A[j] - A[i]
			k, ok := set[hValue]
			if hValue < A[i] && ok {
				cache[i][j] = cache[k][i] + 1
				if max < cache[i][j] {
					max = cache[i][j]
				}
			}
		}
	}

	if max == 0 {
		return 0
	}
	return max + 2
}

func Test873(t *testing.T) {
	assert.Equal(t, 5, lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 3, lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func knightProbability(N int, K int, r int, c int) float64 {
	cache := make([][][]float64, K+1)
	for i := 0; i < K+1; i++ {
		cache[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			cache[i][j] = make([]float64, N)
		}
	}

	return knightProbabilityWithMemo(N, K, r, c, cache)
}

func knightProbabilityWithMemo(N int, K int, r int, c int, cache [][][]float64) float64 {
	if r < 0 || r >= N || c < 0 || c >= N {
		return 0
	}

	if K == 0 {
		return 1
	}

	if cache[K][r][c] != 0 {
		return cache[K][r][c]
	}

	result := knightProbabilityWithMemo(N, K-1, r+1, c+2, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r+2, c+1, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r-1, c+2, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r-2, c+1, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r+1, c-2, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r+2, c-1, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r-1, c-2, cache)*0.125 +
		knightProbabilityWithMemo(N, K-1, r-2, c-1, cache)*0.125

	cache[K][r][c] = result
	return result
}

func Test688(t *testing.T) {
	assert.Equal(t, 0.0625, knightProbability(3, 2, 0, 0))
	assert.Equal(t, 0.0, knightProbability(1, 1, 0, 0))
}

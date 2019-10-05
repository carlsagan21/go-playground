package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	valueToLastIndex := make(map[int]int)
	for i, n := range nums {
		lastIndex, ok := valueToLastIndex[n]
		if !ok {
			valueToLastIndex[n] = i
		} else {
			if i-lastIndex <= k {
				return true
			} else {
				valueToLastIndex[n] = i
			}
		}
	}
	return false
}

func Test219(t *testing.T) {
	assert.Equal(t, true, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	assert.Equal(t, true, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	assert.Equal(t, false, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

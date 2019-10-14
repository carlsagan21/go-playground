package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func singleNumber(nums []int) int {
	return singleNumberBit(nums)
}

func singleNumberMath(nums []int) int {
	sum := 0
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
		sum += n
	}

	uniqSum := 0
	for n := range set {
		uniqSum += n
	}

	return uniqSum*2 - sum
}

func singleNumberBit(nums []int) int {
	// a xor 0 = a
	// a xor a = 0
	// a xor b xor a = a xor a xor b = 0 xor b = b
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func Test136(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
}

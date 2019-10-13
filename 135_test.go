package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// I think it is O(n), but maybe further optimizations possible.
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	if len(ratings) == 1 {
		return 1
	}

	if len(ratings) == 2 {
		if ratings[0] != ratings[1] {
			return 3
		} else {
			return 2
		}
	}

	var higherThan = make([]int, len(ratings))
	// 1-> none, 2->left, 3->right, 4->both
	for i, r := range ratings {
		if i == 0 {
			if r > ratings[i+1] {
				higherThan[i] = 3
			} else {
				higherThan[i] = 1
			}
		} else if i == len(ratings)-1 {
			if r > ratings[i-1] {
				higherThan[i] = 2
			} else {
				higherThan[i] = 1
			}
		} else {
			if r > ratings[i-1] && r > ratings[i+1] {
				higherThan[i] = 4
			} else if r > ratings[i+1] {
				higherThan[i] = 3
			} else if r > ratings[i-1] {
				higherThan[i] = 2
			} else {
				higherThan[i] = 1
			}
		}
	}

	result := memoResult(higherThan)
	sum := 0
	for _, v := range result {
		sum += v
	}
	return sum
}

func memoResult(higherThan []int) []int {
	var cache = make([]int, len(higherThan))
	var result = make([]int, len(higherThan))
	for i := 0; i < len(higherThan); i++ {
		result[i] = getValue(higherThan, i, cache)
	}
	return result
}

func getValue(higherThan []int, i int, cache []int) int {
	if cache[i] != 0 {
		return cache[i]
	}

	if higherThan[i] == 1 {
		cache[i] = 1
		return 1
	} else if higherThan[i] == 2 {
		cache[i] = getValue(higherThan, i-1, cache) + 1
		return cache[i]
	} else if higherThan[i] == 3 {
		cache[i] = getValue(higherThan, i+1, cache) + 1
		return cache[i]
	} else {
		right := getValue(higherThan, i+1, cache)
		left := getValue(higherThan, i-1, cache)
		if right > left {
			cache[i] = right + 1
			return cache[i]
		} else {
			cache[i] = left + 1
			return cache[i]
		}
	}
}

func Test135(t *testing.T) {
	assert.Equal(t, 5, candy([]int{1, 0, 2}))
	assert.Equal(t, 4, candy([]int{1, 2, 2}))
}

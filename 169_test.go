package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// There are 6 ways to solve this problems.
// Boyer-Moore Voting Algorithm https://en.wikipedia.org/wiki/Boyer–Moore_majority_vote_algorithm
// 보이어-무어 투표 알고리즘
func majorityElement(nums []int) int {
	return majorityElementBoyerMooreVoting(nums)
}

func majorityElementBoyerMooreVoting(nums []int) int {
	candidate := nums[0]
	counter := 0
	for _, n := range nums {
		if counter == 0 {
			candidate = n
		}
		if n == candidate {
			counter++
		} else {
			counter--
		}
	}

	return candidate
}

func majorityElementMap(nums []int) int {
	counter := map[int]int{}
	for _, n := range nums {
		count, ok := counter[n]
		if ok {
			counter[n] = count + 1
		} else {
			counter[n] = 1
		}
	}

	for k, v := range counter {
		if len(nums)%2 == 0 {
			if v >= len(nums)/2 {
				return k
			}
		} else {
			if v > len(nums)/2 {
				return k
			}
		}
	}
	return 0
}

func Test169(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

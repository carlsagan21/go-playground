package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isHappy(n int) bool {
	record := make(map[int]struct{})
	return isHappyRecursive(n, record)
}

func isHappyRecursive(curr int, record map[int]struct{}) bool {
	if _, ok := record[curr]; ok {
		return false
	}
	record[curr] = struct{}{}
	next := sumOfSquaresOfDigits(toDigitSlice(curr))
	if next == 1 {
		return true
	} else {
		return isHappyRecursive(next, record)
	}
}

func sumOfSquaresOfDigits(digits []int) int {
	sum := 0
	for _, d := range digits {
		sum += d * d
	}
	return sum
}

func toDigitSlice(n int) []int {
	var slice []int
	for n/10 != 0 {
		slice = append(slice, n%10)
		n /= 10
	}
	return append(slice, n)
}

func Test202(t *testing.T) {
	assert.Equal(t, true, isHappy(19))
	assert.Equal(t, true, isHappy(82))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Passed but time complexity is N^2. There is N solution.
func scoreOfParentheses(S string) int {
	if isUnit(S) {
		return 1
	}
	slicedS := sliceP(S)
	result := 0
	for _, s := range slicedS {
		if isUnit(s) {
			result++
		} else {
			result += 2 * scoreOfParentheses(s[1:len(s)-1])
		}
	}

	return result
}

func sliceP(s string) []string {
	var splitPoints []int
	var stack []bool
	for index, char := range []rune(s) {
		if char == '(' {
			stack = append(stack, true)
		} else {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			splitPoints = append(splitPoints, index)
		}
	}
	var result []string
	lastPoint := 0
	for _, point := range splitPoints {
		result = append(result, s[lastPoint:point+1])
		lastPoint = point + 1
	}
	return result
}

func isUnit(S string) bool {
	if len(S) == 2 {
		return true
	} else {
		return false
	}
}

func Test856(t *testing.T) {
	assert.Equal(t, 1, scoreOfParentheses("()"))
	assert.Equal(t, 2, scoreOfParentheses("(())"))
	assert.Equal(t, 2, scoreOfParentheses("()()"))
	assert.Equal(t, 6, scoreOfParentheses("(()(()))"))
}

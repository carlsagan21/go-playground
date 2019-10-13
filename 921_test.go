package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minAddToMakeValid(S string) int {
	left := 0
	right := 0
	for _, c := range []rune(S) {
		if c == '(' {
			left++
		} else {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left + right
}

func Test921(t *testing.T) {
	assert.Equal(t, 1, minAddToMakeValid("())"))
	assert.Equal(t, 3, minAddToMakeValid("((("))
	assert.Equal(t, 0, minAddToMakeValid("()"))
	assert.Equal(t, 4, minAddToMakeValid("()))(("))
}

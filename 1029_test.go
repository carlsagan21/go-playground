package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// There seems to be better solution.
func removeDuplicates(s string, k int) string {
	dupCount := 0
	var prevChar rune
	var dupLoc []int
	for i, c := range []rune(s) {
		if prevChar == c {
			dupCount++
		} else {
			dupCount = 1
			prevChar = c
		}
		if dupCount == k {
			dupCount = 0
			dupLoc = append(dupLoc, i)
		}
	}
	reverse(dupLoc)

	if len(dupLoc) == 0 {
		return s
	}

	for _, loc := range dupLoc {
		s = s[0:loc+1-k] + s[loc+1:]
	}
	return removeDuplicates(s, k)
}

func reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func Test1029(t *testing.T) {
	assert.Equal(t, "abcd", removeDuplicates("abcd", 2))
	assert.Equal(t, "aa", removeDuplicates("deeedbbcccbdaa", 3))
	assert.Equal(t, "ps", removeDuplicates("pbbcggttciiippooaais", 2))
}

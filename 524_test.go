package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findLongestWord(s string, d []string) string {
	index := make([]int, len(d))
	var result string
	for _, c := range []byte(s) {
		for i, w := range d {
			if index[i] >= len(w) {
				continue
			}
			if w[index[i]] == c {
				index[i]++
				if index[i] >= len(w) {
					if len(result) == len(w) && result > w {
						result = w
					}
					if len(result) < len(w) {
						result = w
					}
				}
			}
		}
	}

	return result
}

func Test524(t *testing.T) {
	assert.Equal(t, "apple", findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	assert.Equal(t, "a", findLongestWord("abpcplea", []string{"a", "b", "c"}))
}

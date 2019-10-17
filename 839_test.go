package main

import (
	"github.com/carlsagan21/go-playground/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func numSimilarGroups(A []string) int {
	disjointSet := lib.NewDisjointSet(len(A))
	for i, wi := range A {
		for j := i + 1; j < len(A); j++ {
			wj := A[j]
			if similar(wi, wj) {
				disjointSet.Union(i, j)
			}
		}
	}
	return disjointSet.Count()
}

func similar(a string, b string) bool {
	ab := []byte(a)
	bb := []byte(b)
	diffCount := 0
	for i := 0; i < len(a); i++ {
		if ab[i] != bb[i] {
			diffCount++
		}
		if diffCount > 2 {
			return false
		}
	}
	return true
}

func Test839(t *testing.T) {
	assert.Equal(t, 2, numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
	assert.Equal(t, 1, numSimilarGroups([]string{"bpkvu", "bukvp", "vpkbu", "vpkub"}))
	assert.Equal(t, 1, numSimilarGroups([]string{"jmijc", "imjjc", "jcijm", "cmijj", "mijjc"}))
}

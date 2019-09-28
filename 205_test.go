package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test205(t *testing.T) {
	assert.Equal(t, true, isIsomorphic("egg", "add"))
	assert.Equal(t, false, isIsomorphic("foo", "bar"))
	assert.Equal(t, true, isIsomorphic("paper", "title"))
}

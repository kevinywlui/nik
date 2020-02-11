package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	assert := assert.New(t)
	assertIsSubSequence := func(small string, big string) bool {
		return assert.True(IsSubsequence(small, big))
	}
	assertNotIsSubSequence := func(small string, big string) bool {
		return assert.False(IsSubsequence(small, big))
	}
	assertIsSubSequence("abc", "abcd")
	assertIsSubSequence("", "")
	assertIsSubSequence("a", "a")
	assertIsSubSequence("ad", "abcd")
	assertIsSubSequence("z", "kevinkevinkevinz")

	assertNotIsSubSequence("abc", "xyz")
	assertNotIsSubSequence("abc", "kevinkevinkevin")
	assertNotIsSubSequence("a", "")
}

func TestIsBaseSubsequence(t *testing.T) {
	assert := assert.New(t)
	assertIsBaseSubsequence := func(small string, big string) bool {
		return assert.True(IsBaseSubsequence(small, big))
	}
	assertNotIsBaseSubsequence := func(small string, big string) bool {
		return assert.False(IsBaseSubsequence(small, big))
	}
	assertNotIsBaseSubsequence("abc", "/abc/xyz")
	assertNotIsBaseSubsequence("abc", "/a/bc")
	assertIsBaseSubsequence("abc", "/abc")
	assertIsBaseSubsequence("abc", "/abc/")
}

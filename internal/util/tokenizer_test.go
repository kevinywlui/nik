package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixBaseSplit(t *testing.T) {
	assert := assert.New(t)
	assertPrefixBaseSplit := func(path, prefix, base string) bool {
		tokens := PrefixBaseSplit(path)
		got_prefix := tokens[0]
		got_base := tokens[1]
		return assert.Equal(got_prefix, prefix) && assert.Equal(got_base, base)
	}
	assertPrefixBaseSplit("/abc/xyz", "/abc/", "xyz")
	assertPrefixBaseSplit("/abc/xyz/", "/abc/", "xyz")
	assertPrefixBaseSplit("/xyz", "/", "xyz")
	assertPrefixBaseSplit("/", "/", "")
}

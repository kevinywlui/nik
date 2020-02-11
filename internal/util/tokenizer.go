package util

import (
	"strings"
)

func PrefixBaseSplit(path string) [2]string {
	path_split := strings.Split(path, "/")
	end_slash := (path[len(path)-1] == '/')
	var base_idx int
	if end_slash {
		base_idx = len(path_split) - 2
	} else {
		base_idx = len(path_split) - 1
	}
	base := path_split[base_idx]
	prefix := strings.Join(path_split[:base_idx], "/")
	prefix += "/"

	tokens := [2]string{prefix, base}
	return tokens
}

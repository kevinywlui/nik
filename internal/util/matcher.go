package util

import (
	"path/filepath"
)

// IsSubsequence will determine if `small` is a subsequence of `big`.
func IsSubsequence(small string, big string) bool {

	// Handle the empty string case
	if big == "" {
		return small == ""
	}

	idx_b := 0
	for idx_s := 0; idx_s < len(small); {
		if small[idx_s] == big[idx_b] {
			idx_s++
			idx_b++
		} else {
			idx_b++
		}
		// Return an false if we reach the end of `big` but not `small`
		if (idx_b == len(big)) && (idx_s < len(small)) {
			return false
		}
	}
	return true
}

// IsBaseSubsequence will determine if `small` is a subsequence of the filepath
// base of `big`.
func IsBaseSubsequence(small string, big string) bool {
	big_base := filepath.Base(big)
	return IsSubsequence(small, big_base)
}

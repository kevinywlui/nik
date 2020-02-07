package util

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type matcher_t func(string, string) bool
type asserter_t func(string, string, bool)

func boolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func getFunctionName(f interface{}) string {
	f_ptr := reflect.ValueOf(f).Pointer()
	full_name := runtime.FuncForPC(f_ptr).Name()
	split_full_name := strings.Split(full_name, "/")
	return split_full_name[len(split_full_name)-1]
}

func asserterBuilder(t *testing.T, f matcher_t) asserter_t {
	return func(small string, big string, expect bool) {
		got := f(small, big)
		if got != expect {
			f_name := getFunctionName(f)
			expect_str := boolToString(expect)
			got_str := boolToString(got)
			t.Errorf("Got %s(%q, %q)==%s, expected %s", f_name, small, big, got_str, expect_str)
		}
	}
}

func TestIsSubsequence(t *testing.T) {
	asserter := asserterBuilder(t, IsSubsequence)

	asserter("abc", "abcd", true)
	asserter("abc", "xyz", false)
	asserter("", "", true)
	asserter("a", "a", true)
	asserter("a", "", false)
	asserter("ad", "abcd", true)
	asserter("abc", "kevinkevinkevin", false)
	asserter("z", "kevinkevinkevinz", true)
}

func TestIsBaseSubsequence(t *testing.T) {
	asserter := asserterBuilder(t, IsBaseSubsequence)

	asserter("abc", "/abc/xyz", false)
	asserter("abc", "/a/bc", false)
	asserter("abc", "/abc", true)
	asserter("abc", "/abc/", true)
}

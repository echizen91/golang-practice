package easy

import (
	"reflect"
	"strings"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.4 MB, less than 7.75% of Go online submissions for Longest Common Prefix.
*/

// LongestCommonPrefix : find longest common prefix amongst all strings in the array
func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var result string
	// find the shortest string
	shortS := strs[0]
	for _, v := range strs[1:] {
		if len(shortS) > len(v) {
			shortS = v
		}
	}

	for i := 0; i < len(shortS); i++ {
		b := true
		result = string(shortS[0 : len(shortS)-i])
		for _, v := range strs {
			if reflect.DeepEqual(shortS, v) {
				continue
			}
			if !strings.HasPrefix(v, result) {
				b = false
				result = ""
				break
			}
		}
		if b {
			break
		}
	}

	return result
}

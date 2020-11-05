package medium

import (
	"math"
	"reflect"
	"strings"
)

/*
Given a string s, find the length of the longest substring without repeating characters.

First Attempt Solution:
Runtime: 8 ms, faster than 73.29% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 6.1 MB, less than 9.48% of Go online submissions for Longest Substring Without Repeating Characters.
*/

func getIndex(s string, c string) int {
	for i, v := range s {
		if reflect.DeepEqual(string(v), c) {
			return i
		}
	}
	// won't happen
	return -1
}

func lengthOfLongestSubstring(s string) int {
	var current, global int
	var currentString string

	// First Attempt
	for _, v := range s {
		// if current evaluating char not in currentString
		if strings.Contains(currentString, string(v)) {
			index := getIndex(currentString, string(v))
			// log.Println(index)
			currentString = currentString[index+1:] + string(v)
			current = len(currentString)
			global = int(math.Max(float64(global), float64(current)))
		} else {
			currentString += string(v)
			current++
			global = int(math.Max(float64(global), float64(current)))
		}
	}

	// log.Println(current, global)
	return global
}

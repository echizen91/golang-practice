package easy

import "strings"

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
Memory Usage: 2.3 MB, less than 5.06% of Go online submissions for Implement strStr().
*/

// StrStr : finding the needle in the haystack
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) == 0 && len(haystack) > 0 {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		substack := haystack[i:]
		if strings.HasPrefix(substack, needle) {
			return i
		}
	}
	return -1
}

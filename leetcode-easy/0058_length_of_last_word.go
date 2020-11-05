package easy

import "strings"

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
return the length of last word (last word means the last appearing word if we loop from left to right) in the string.

If the last word does not exist, return 0.

Note: A word is defined as a maximal substring consisting of non-space characters only.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Length of Last Word.
*/

// LengthOfLastWord : Return length of last word, else 0 if it doesn't exist
func LengthOfLastWord(s string) int {
	if len(s) < 1 {
		return 0
	}
	sArr := strings.Split(strings.Trim(s, " "), " ")

	return len(sArr[len(sArr)-1])
}

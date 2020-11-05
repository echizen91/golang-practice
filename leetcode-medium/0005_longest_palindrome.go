package medium

/*
Given a string s, return the longest palindromic substring in s.

Note: Test case only takes into account the first longest string
*/

func longestPalindrome(s string) string {
	ll := len(s)
	if ll == 0 {
		return ""
	}

	var l, r, pl, pr int
	for r < ll {
		// gobble up dup chars
		for r+1 < ll && s[l] == s[r+1] {
			// log.Println("gobble", string(s[l]), string(s[r+1]))
			r++
		}
		// find size of this palindrome
		for l-1 >= 0 && r+1 < ll && s[l-1] == s[r+1] {
			// log.Println("find size of palindrome", string(s[l-1]), string(s[r+1]))
			l--
			r++
		}
		if r-l > pr-pl {
			pl, pr = l, r
		}
		// reset to next mid point
		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}

package hard

import "strings"

/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
*/

func isMatch(s string, p string) bool {
	var wildcard int
	// wildcardChar := make(map[string]string, 0)

	for i := 0; i < len(p); i++ {
		// get current char
		current := string(p[i])

		// check next char if it's *
		var match bool
		if i+1 < len(p) && string(p[i+1]) == "*" {
			match = true
			i++
		}

		// if wildcard match
		if current == "." {
			// with *, definitely true
			if match {
				wildcard = 20
				continue
			}
			// else increase wildcard count
			wildcard++
			continue
		}

		// if no *, check for only 1 char
		if !match {
			index := strings.Index(s, current)
			if index > -1 {
				s = strings.Replace(s, current, "", 1)
			} else {
				return false
			}
		} else {
			index := strings.Index(s, current)
			if index > -1 {
				for j := index; j < len(s); {
					if i+1 < len(p) && string(p[i+1]) == current {
						i++
					}
					if string(s[j]) == current {
						s = strings.Replace(s, current, "", 1)
					} else {
						break
					}
				}
			}
		}
	}

	// for _, c := range wildcardChar {
	// 	s = strings.ReplaceAll(s, c, "")
	// }

	if len(s)-wildcard <= 0 {
		return true
	}

	return false
}

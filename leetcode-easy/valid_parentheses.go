package easy

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2.1 MB, less than 5.83% of Go online submissions for Valid Parentheses.
*/

// ValidParentheses : determine valid string with parentheses
func ValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack []string
	result := true
	for _, c := range s {
		if string(c) == "(" {
			stack = append(stack, "(")
			continue
		}
		if string(c) == "{" {
			stack = append(stack, "{")
			continue
		}
		if string(c) == "[" {
			stack = append(stack, "[")
			continue
		}
		if len(stack) > 0 {
			if string(c) == ")" && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				continue
			} else if string(c) == "}" && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
				continue
			} else if string(c) == "]" && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
				continue
			} else {
				result = false
				break
			}
		} else {
			result = false
			break
		}

	}

	if len(stack) > 0 {
		result = false
	}

	return result
}

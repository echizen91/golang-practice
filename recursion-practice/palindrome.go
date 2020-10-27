package recursion

import "reflect"

/*
Recursion to determine whether string is a palindrome
*/

func palindromeCheck(str string) bool {
	if len(str) < 2 {
		return true
	}
	if reflect.DeepEqual(str[0], str[len(str)-1]) == true {
		return palindromeCheck(str[1:len(str)-1]) && true
	}
	return false
}

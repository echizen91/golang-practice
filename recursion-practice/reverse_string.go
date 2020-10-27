package recursion

/*
Reverse a string using recursion
*/

func recursion(str string) string {
	if len(str) < 2 {
		return str
	}
	return recursion(str[1:]) + string(str[0])
}

func reverseString(str string) string {
	return recursion(str)
}

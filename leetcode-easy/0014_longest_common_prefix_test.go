package easy

/*
Input: strs = ["flower","flow","flight"]
Output: "fl"

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:
0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.
*/

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	e1 := []string{
		"flower", "flow", "flight",
	}
	e1Expected := "fl"
	e1Result := LongestCommonPrefix(e1)

	if reflect.DeepEqual(e1Result, e1Expected) != true {
		t.Errorf("Longest Common Prefix Incorrect: %v ; wants: %v", e1Result, e1Expected)
	}

	e2 := []string{
		"dog", "racecar", "car",
	}
	e2Expected := ""
	e2Result := LongestCommonPrefix(e2)

	if reflect.DeepEqual(e2Result, e2Expected) != true {
		t.Errorf("Longest Common Prefix Incorrect: %v ; wants: %v", e2Result, e2Expected)
	}

	e3 := []string{
		"ab", "a",
	}
	e3Expected := "a"
	e3Result := LongestCommonPrefix(e3)

	if reflect.DeepEqual(e3Result, e3Expected) != true {
		t.Errorf("Longest Common Prefix Incorrect: %v ; wants: %v", e3Result, e3Expected)
	}

	e4 := []string{
		"a",
	}
	e4Expected := "a"
	e4Result := LongestCommonPrefix(e4)

	if reflect.DeepEqual(e4Result, e4Expected) != true {
		t.Errorf("Longest Common Prefix Incorrect: %v ; wants: %v", e4Result, e4Expected)
	}

	e5 := []string{}
	e5Expected := ""
	e5Result := LongestCommonPrefix(e5)

	if reflect.DeepEqual(e5Result, e5Expected) != true {
		t.Errorf("Longest Common Prefix Incorrect: %v ; wants: %v", e5Result, e5Expected)
	}
}

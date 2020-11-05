package easy

/*
Input: s = "III"
Output: 3

Input: s = "IV"
Output: 4

Input: s = "IX"
Output: 9

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Constraints:
1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	e1 := "III"
	e1Expected := 3
	e1Result := RomanToInt(e1)

	if reflect.DeepEqual(e1Result, e1Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e1Result, e1Expected)
	}

	e2 := "IV"
	e2Expected := 4
	e2Result := RomanToInt(e2)

	if reflect.DeepEqual(e2Result, e2Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e2Result, e2Expected)
	}

	e3 := "IX"
	e3Expected := 9
	e3Result := RomanToInt(e3)

	if reflect.DeepEqual(e3Result, e3Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e3Result, e3Expected)
	}

	e4 := "LVIII"
	e4Expected := 58
	e4Result := RomanToInt(e4)

	if reflect.DeepEqual(e4Result, e4Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e4Result, e4Expected)
	}

	e5 := "MCMXCIV"
	e5Expected := 1994
	e5Result := RomanToInt(e5)

	if reflect.DeepEqual(e5Result, e5Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e5Result, e5Expected)
	}

	e6 := "CDXL"
	e6Expected := 440
	e6Result := RomanToInt(e6)

	if reflect.DeepEqual(e6Result, e6Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e6Result, e6Expected)
	}

	e7 := "DCX"
	e7Expected := 610
	e7Result := RomanToInt(e7)

	if reflect.DeepEqual(e7Result, e7Expected) != true {
		t.Errorf("Roman To Int Incorrect: %v ; wants: %v", e7Result, e7Expected)
	}
}

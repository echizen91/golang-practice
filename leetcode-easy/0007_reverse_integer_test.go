package easy

import (
	"reflect"
	"testing"
)

/*
Input: x = 123
Output: 321

Input: x = -123
Output: -321

Input: x = 120
Output: 21

Input: x = 0
Output: 0
*/

func TestReverseInteger(t *testing.T) {
	e1 := 123
	e1Expected := 321
	e1 = ReverseInteger(e1)

	if reflect.DeepEqual(e1, e1Expected) != true {
		t.Errorf("Reverse Integer Incorrect: %v ; wants: %v", e1, e1Expected)
	}

	e2 := -123
	e2Expected := -321
	e2 = ReverseInteger(e2)

	if reflect.DeepEqual(e2, e2Expected) != true {
		t.Errorf("Reverse Integer Incorrect: %v ; wants: %v", e2, e2Expected)
	}

	e3 := 120
	e3Expected := 21
	e3 = ReverseInteger(e3)

	if reflect.DeepEqual(e3, e3Expected) != true {
		t.Errorf("Reverse Integer Incorrect: %v ; wants: %v", e3, e3Expected)
	}

	e4 := 0
	e4Expected := 0
	e4 = ReverseInteger(e4)

	if reflect.DeepEqual(e4, e4Expected) != true {
		t.Errorf("Reverse Integer Incorrect: %v ; wants: %v", e4, e4Expected)
	}

	e5 := 1534236469
	e5Expected := 0
	e5 = ReverseInteger(e5)

	if reflect.DeepEqual(e5, e5Expected) != true {
		t.Errorf("Reverse Integer Incorrect: %v ; wants: %v", e5, e5Expected)
	}
}

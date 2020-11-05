package easy

import (
	"reflect"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	e1 := 121
	e1Expected := true
	e1Result := PalindromeNumber(e1)

	if reflect.DeepEqual(e1Result, e1Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e1Result, e1Expected)
	}

	e2 := -121
	e2Expected := false
	e2Result := PalindromeNumber(e2)

	if reflect.DeepEqual(e2Result, e2Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e2Result, e2Expected)
	}

	e3 := 10
	e3Expected := false
	e3Result := PalindromeNumber(e3)

	if reflect.DeepEqual(e3Result, e3Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e3Result, e3Expected)
	}

	e4 := -101
	e4Expected := false
	e4Result := PalindromeNumber(e4)

	if reflect.DeepEqual(e4Result, e4Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e4Result, e4Expected)
	}

	e5 := 12321
	e5Expected := true
	e5Result := PalindromeNumber(e5)

	if reflect.DeepEqual(e5Result, e5Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e5Result, e5Expected)
	}
}

func TestPalindromeNumberOptimized(t *testing.T) {
	e1 := 121
	e1Expected := true
	e1Result := PalindromeNumberOptimized(e1)

	if reflect.DeepEqual(e1Result, e1Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e1Result, e1Expected)
	}

	e2 := -121
	e2Expected := false
	e2Result := PalindromeNumberOptimized(e2)

	if reflect.DeepEqual(e2Result, e2Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e2Result, e2Expected)
	}

	e3 := 10
	e3Expected := false
	e3Result := PalindromeNumberOptimized(e3)

	if reflect.DeepEqual(e3Result, e3Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e3Result, e3Expected)
	}

	e4 := -101
	e4Expected := false
	e4Result := PalindromeNumberOptimized(e4)

	if reflect.DeepEqual(e4Result, e4Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e4Result, e4Expected)
	}

	e5 := 12321
	e5Expected := true
	e5Result := PalindromeNumberOptimized(e5)

	if reflect.DeepEqual(e5Result, e5Expected) != true {
		t.Errorf("Palindrome Number Incorrect: %v ; wants: %v", e5Result, e5Expected)
	}
}

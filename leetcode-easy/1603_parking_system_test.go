package easy

import (
	"reflect"
	"testing"
)

func TestConstructorParking(t *testing.T) {
	p := ConstructorParking(1, 1, 0)

	f1 := p.AddCar(1)
	expectedF1 := true

	if reflect.DeepEqual(f1, expectedF1) != true {
		t.Errorf("Add Big Car Error. got: %v, want: %v", f1, expectedF1)
	}

	f2 := p.AddCar(2)
	expectedF2 := true

	if reflect.DeepEqual(f2, expectedF2) != true {
		t.Errorf("Add Big Car Error. got: %v, want: %v", f2, expectedF2)
	}

	f3 := p.AddCar(3)
	expectedF3 := false

	if reflect.DeepEqual(f3, expectedF3) != true {
		t.Errorf("Add Big Car Error. got: %v, want: %v", f3, expectedF3)
	}

	f4 := p.AddCar(1)
	expectedF4 := false

	if reflect.DeepEqual(f4, expectedF4) != true {
		t.Errorf("Add Big Car Error. got: %v, want: %v", f4, expectedF4)
	}
}

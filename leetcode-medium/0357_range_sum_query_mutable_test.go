package medium

import (
	"reflect"
	"testing"
)

func TestNumArray(t *testing.T) {
	// check for object creation
	obj := Constructor([]int{1, 3, 5})

	// test values
	if reflect.DeepEqual(obj.nums, []int{1, 3, 5}) != true {
		t.Errorf("object value creation incorrect, want: %v; got: %v",
			[]int{1, 3, 5}, obj.nums)
	}

	// test type
	if reflect.TypeOf(obj) != reflect.TypeOf(NumArray{}) {
		t.Errorf("object type different, want: %v; got: %v",
			reflect.TypeOf(NumArray{}), reflect.TypeOf(obj))
	}

	// test sum range
	val := obj.SumRange(0, 2)
	res := 9

	if reflect.DeepEqual(val, res) != true {
		t.Errorf("sum range (1) incorrect, want: %v; got: %v",
			res, val)
	}

	// test update num array & sum again
	obj.Update(1, 2)
	val = obj.SumRange(0, 2)
	res = 8

	if reflect.DeepEqual(val, res) != true {
		t.Errorf("sum range (2) incorrect, want: %v; got: %v",
			res, val)
	}
}

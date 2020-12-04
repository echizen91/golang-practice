package medium

import (
	"reflect"
	"testing"
)

func TestNumMatrix(t *testing.T) {
	obj := ConstructorMatrix(
		[][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		})

	testObj := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	// test value of matrix
	if reflect.DeepEqual(obj.numMatrix, testObj) != true {
		t.Errorf("object value creation incorrect, want: %v; got: %v",
			testObj, obj.numMatrix)
	}

	// test type
	if reflect.TypeOf(obj) != reflect.TypeOf(NumMatrix{}) {
		t.Errorf("object type different, want: %v; got: %v",
			reflect.TypeOf(NumMatrix{}), reflect.TypeOf(obj))
	}

	// test sum region (1)
	val := obj.SumRegion(2, 1, 4, 3)
	res := 8

	if reflect.DeepEqual(val, res) != true {
		t.Errorf("sum range (1) incorrect, want: %v; got: %v",
			res, val)
	}

	// test sum region (2)
	val = obj.SumRegion(1, 1, 2, 2)
	res = 11

	if reflect.DeepEqual(val, res) != true {
		t.Errorf("sum range (2) incorrect, want: %v; got: %v",
			res, val)
	}

	// test sum region
	val = obj.SumRegion(1, 2, 2, 4)
	res = 12

	if reflect.DeepEqual(val, res) != true {
		t.Errorf("sum range (3) incorrect, want: %v; got: %v",
			res, val)
	}
}

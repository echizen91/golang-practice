package easy

import (
	"reflect"
	"testing"
)

func TestTwoSumsExampleOne(t *testing.T) {
	nums := []int{
		2, 7, 11, 15,
	}

	target := 9

	result := TwoSums(nums, target)
	expected := []int{
		0, 1,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Incorrect: %v ; wants: %v", result, expected)
	}
}

func TestTwoSumsExampleTwo(t *testing.T) {
	nums := []int{
		3, 2, 4,
	}

	target := 6

	result := TwoSums(nums, target)
	expected := []int{
		1, 2,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Incorrect: %v ; wants: %v", result, expected)
	}
}

func TestTwoSumsExampleThree(t *testing.T) {
	nums := []int{
		3, 3,
	}

	target := 6

	result := TwoSums(nums, target)
	expected := []int{
		0, 1,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Incorrect: %v ; wants: %v", result, expected)
	}
}

func TestTwoSumsOptimizedExampleOne(t *testing.T) {
	nums := []int{
		2, 7, 11, 15,
	}

	target := 9

	result := TwoSumsOptimized(nums, target)
	expected := []int{
		0, 1,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Optimized Incorrect: %v ; wants: %v", result, expected)
	}
}

func TestTwoSumsOptimizedExampleTwo(t *testing.T) {
	nums := []int{
		3, 2, 4,
	}

	target := 6

	result := TwoSumsOptimized(nums, target)
	expected := []int{
		1, 2,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Optimized Incorrect: %v ; wants: %v", result, expected)
	}
}

func TestTwoSumsOptimizedExampleThree(t *testing.T) {
	nums := []int{
		3, 3,
	}

	target := 6

	result := TwoSumsOptimized(nums, target)
	expected := []int{
		0, 1,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Optimized Incorrect: %v ; wants: %v", result, expected)
	}
}

func TestTwoSumsOptimizedNoMatch(t *testing.T) {
	nums := []int{
		3, 3,
	}

	target := 7

	result := TwoSumsOptimized(nums, target)
	expected := []int{}

	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Two Sums Optimized Incorrect: %v ; wants: %v", result, expected)
	}
}

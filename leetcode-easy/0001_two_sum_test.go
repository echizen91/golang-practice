package easy

/*
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

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

package matrix

import (
	"reflect"
	"testing"
)

func TestRotateMatrixRightFaster(t *testing.T) {
	m := GetTwoByTwo()
	m.RotateMatrixRightFaster()

	compareM := &Matrix{
		matrix: [][]int{
			{2, 4},
			{1, 3},
		},
	}

	if reflect.DeepEqual(m, compareM) != true {
		t.Errorf("Rotate Matrix Right Incorrect: %v ; wants: %v", m, compareM)
	}

	m2 := GetThreeByThree()
	m2.RotateMatrixRightFaster()

	compareM2 := &Matrix{
		matrix: [][]int{
			{3, 6, 9},
			{2, 5, 8},
			{1, 4, 7},
		},
	}

	if reflect.DeepEqual(m2, compareM2) != true {
		t.Errorf("Rotate Matrix Right Incorrect: %v ; wants: %v", m2, compareM2)
	}

	m3 := GetFourByFour()
	m3.RotateMatrixRightFaster()

	compareM3 := &Matrix{
		matrix: [][]int{
			{4, 8, 12, 16},
			{3, 7, 11, 15},
			{2, 6, 10, 14},
			{1, 5, 9, 13},
		},
	}

	if reflect.DeepEqual(m3, compareM3) != true {
		t.Errorf("Rotate Matrix Right Incorrect: %v ; wants: %v", m3, compareM3)
	}
}

func TestMatrix_RotateMatrixLeftFaster(t *testing.T) {
	m := GetTwoByTwo()
	m.RotateMatrixLeftFaster()

	compareM := &Matrix{
		matrix: [][]int{
			{3, 1},
			{4, 2},
		},
	}

	if reflect.DeepEqual(m, compareM) != true {
		t.Errorf("Rotate Matrix Left Incorrect: %v ; wants: %v", m, compareM)
	}

	m2 := GetThreeByThree()
	m2.RotateMatrixLeftFaster()

	compareM2 := &Matrix{
		matrix: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
	}

	if reflect.DeepEqual(m2, compareM2) != true {
		t.Errorf("Rotate Matrix Left Incorrect: %v ; wants: %v", m2, compareM2)
	}

	m3 := GetFourByFour()
	m3.RotateMatrixLeftFaster()

	compareM3 := &Matrix{
		matrix: [][]int{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		},
	}

	if reflect.DeepEqual(m3, compareM3) != true {
		t.Errorf("Rotate Matrix Left Incorrect: %v ; wants: %v", m3, compareM3)
	}
}

package medium

// NumMatrix : structure of matrix
type NumMatrix struct {
	numMatrix [][]int
}

// ConstructorMatrix : returns a creation of NumMatrix
func ConstructorMatrix(matrix [][]int) NumMatrix {
	return NumMatrix{
		numMatrix: matrix,
	}
}

// SumRegion : return the sum of numbers within a given 2D range
func (a *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var total int

	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			total += a.numMatrix[i][j]
		}
	}

	return total
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

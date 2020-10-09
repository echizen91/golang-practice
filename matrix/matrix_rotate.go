package matrix

/*
given an nxn matrix, write a method to rotate.
*/

// Matrix : n by n matrix
type Matrix struct {
	matrix [][]int
}

// mxmCreation : create n x n empty matrix
func mxmCreation(size int) [][]int {
	resultMatrix := make([][]int, size)
	for i := range resultMatrix {
		resultMatrix[i] = make([]int, size)
	}
	return resultMatrix
}

// RotateMatrixRight : function to rotate matrix (nxn) 90 degrees to the right -> Worst case: O(n^2)
func (m *Matrix) RotateMatrixRight() {
	resultMatrix := mxmCreation(len(m.matrix))

	for i := 0; i < len(m.matrix); i++ { // e.g. 0 - 2 -> 3x3
		counter := len(m.matrix) - 1
		for j := 0; j < len(m.matrix); j++ {
			resultMatrix[counter][i] = m.matrix[i][j]
			counter--
		}
	}

	m.matrix = resultMatrix
}

// RotateMatrixLeft : function to rotate matrix (nxn) 90 degrees to the left -> Worst case: O(n^2)
func (m *Matrix) RotateMatrixLeft() {
	resultMatrix := mxmCreation(len(m.matrix))

	counter := len(m.matrix) - 1
	for i := 0; i < len(m.matrix); i++ { // e.g. 0 - 2 -> 3x3

		for j := 0; j < len(m.matrix); j++ {
			resultMatrix[j][counter] = m.matrix[i][j]
			// counter--
		}
		counter--
	}

	m.matrix = resultMatrix
}

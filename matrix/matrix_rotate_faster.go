package matrix

// RotateMatrixRightFaster : function to rotate matrix (nxn) 90 degrees to the right -> Worst case: O(n^2 / 2)
func (m *Matrix) RotateMatrixRightFaster() {
	// if 3x3 -> 3 / 2 = 1
	// if 4x4 -> 4 / 2 = 2
	levels := len(m.matrix) / 2

	matrixLen := len(m.matrix) - 1
	// log.Printf("For  %v x %v matrix", matrixLen, matrixLen)
	// Outer loop to cover levels
	for i := 0; i < levels; i++ {
		// Inner loop cover fields to be swapped
		for j := i; j < matrixLen-i; j++ {
			// log.Printf("Swapping (1) [%v,%v] with [%v,%v]", j, matrixLen-i, matrixLen-j, i)
			m.matrix[j][matrixLen-i], m.matrix[matrixLen-j][i] = m.matrix[matrixLen-j][i], m.matrix[j][matrixLen-i]
			// log.Printf("Swapping (2) [%v,%v] with [%v,%v]", i, j, matrixLen-j, i)
			m.matrix[i][j], m.matrix[matrixLen-j][i] = m.matrix[matrixLen-j][i], m.matrix[i][j]
			// log.Printf("Swapping (3) [%v,%v] with [%v,%v]", j, matrixLen-i, matrixLen-i, matrixLen-j)
			m.matrix[j][matrixLen-i], m.matrix[matrixLen-i][matrixLen-j] = m.matrix[matrixLen-i][matrixLen-j], m.matrix[j][matrixLen-i]
		}
		// for _, v := range m.matrix {
		// 	log.Println(v)
		// }
	}
}

// RotateMatrixLeftFaster : function to rotate matrix (nxn) 90 degrees to the left -> Worst case: O(n^2 / 2)
func (m *Matrix) RotateMatrixLeftFaster() {
	levels := len(m.matrix) / 2

	matrixLen := len(m.matrix) - 1

	// Outer loop to cover levels
	for i := 0; i < levels; i++ {
		// Inner loop cover fields to be swapped
		for j := i; j < matrixLen-i; j++ {
			// log.Printf("Swapping (1) [%v,%v] with [%v,%v]", i, j, matrixLen-i, matrixLen-j)
			m.matrix[i][j], m.matrix[matrixLen-i][matrixLen-j] = m.matrix[matrixLen-i][matrixLen-j], m.matrix[i][j]
			// log.Printf("Swapping (2) [%v,%v] with [%v,%v]", i, j, matrixLen-j, i)
			m.matrix[i][j], m.matrix[matrixLen-j][i] = m.matrix[matrixLen-j][i], m.matrix[i][j]
			// log.Printf("Swapping (3) [%v,%v] with [%v,%v]", j, matrixLen-i, matrixLen-i, matrixLen-j)
			m.matrix[j][matrixLen-i], m.matrix[matrixLen-i][matrixLen-j] = m.matrix[matrixLen-i][matrixLen-j], m.matrix[j][matrixLen-i]
		}
		// for _, v := range m.matrix {
		// 	log.Println(v)
		// }
	}
}

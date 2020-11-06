package medium

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

func isValidSudoku(board [][]byte) bool {
	// Check horizonal
	for i := 0; i < len(board); i++ {
		validMap := make(map[byte]byte)
		for j := 0; j < len(board); j++ {
			byteS := board[i][j]
			if byteS == '.' {
				continue
			}
			if _, ok := validMap[byteS]; ok {
				return false
			}
			validMap[byteS] = byteS
		}
	}

	// Check vertical
	for i := 0; i < len(board); i++ {
		validMap := make(map[byte]byte)
		for j := 0; j < len(board); j++ {
			byteS := board[j][i]
			if byteS == '.' {
				continue
			}
			if _, ok := validMap[byteS]; ok {
				return false
			}
			validMap[byteS] = byteS
		}
	}

	/*
		[0-2][0-2]  [0-2][3-5]  [0-2][6-8]
		[3-5][0-2]  [3-5][3-5]  [3-5][6-8]
		[6-8][0-2]	[6-8][3-5]  [6-8][6-8]
	*/
	// Check 3x3 grids
	for i := 0; i < len(board); i += 3 {

		// first box
		validMap := make(map[byte]byte)
		for j := 0; j < 3; j++ {
			for k := i; k < i+3; k++ {
				byteS := board[k][j]
				if byteS == '.' {
					continue
				}
				if _, ok := validMap[byteS]; ok {
					return false
				}
				validMap[byteS] = byteS
			}
		}

		// second box
		validMap = make(map[byte]byte)
		for j := 3; j < 6; j++ {
			for k := i; k < i+3; k++ {
				byteS := board[k][j]
				if byteS == '.' {
					continue
				}
				if _, ok := validMap[byteS]; ok {
					return false
				}
				validMap[byteS] = byteS
			}
		}

		// third box
		validMap = make(map[byte]byte)
		for j := 6; j < 9; j++ {
			for k := i; k < i+3; k++ {
				byteS := board[k][j]
				if byteS == '.' {
					continue
				}
				if _, ok := validMap[byteS]; ok {
					return false
				}
				validMap[byteS] = byteS
			}
		}
	}

	return true
}

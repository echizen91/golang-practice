package easy

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	for i := range rows {
		switch i {
		case 0:
			rows[i] = []int{1}
		case 1:
			rows[i] = []int{1, 1}
		default:
			rows[i] = make([]int, i+1)                 // make a row of size i+1
			rows[i][0], rows[i][len(rows[i])-1] = 1, 1 // set first and last elements to 1
			for j := 1; j <= len(rows[i])-2; j++ {     // calculate rest based on previous row
				rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
			}
		}
	}
	return rows
}

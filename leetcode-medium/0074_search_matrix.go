package medium

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	row, col := len(matrix), len(matrix[0])
	low, high := 0, row*col-1
	for low <= high {
		mid := low + (high-low)/2
		if matrix[mid/col][mid%col] > target {
			high = mid - 1
		} else if matrix[mid/col][mid%col] < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}

package easy

func insertionSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		for j := i; j > 0; j-- {
			if (*arr)[j] < (*arr)[j-1] {
				(*arr)[j], (*arr)[j-1] = (*arr)[j-1], (*arr)[j]
			} else {
				break
			}
		}
	}
}

func canMakeArithmeticProgression(arr []int) bool {

	if len(arr) == 2 {
		return true
	}

	insertionSort(&arr)

	progression := arr[0] - arr[1]
	for i := 1; i < len(arr)-1; i++ {
		cur := arr[i] - arr[i+1]
		if cur != progression {
			return false
		}
	}

	return true
}

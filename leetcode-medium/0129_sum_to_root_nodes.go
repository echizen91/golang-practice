package medium

func recursiveSum(current *TreeNode, val int, total *int) {
	if current == nil {
		*total += val
		return
	}
	val = val*10 + current.Val
	if current.Left == nil && current.Right == nil {
		recursiveSum(current.Left, val, total)
		return
	}
	if current.Left != nil {
		recursiveSum(current.Left, val, total)
	}
	if current.Right != nil {
		recursiveSum(current.Right, val, total)
	}

}

func sumNumbers(root *TreeNode) int {
	var result int
	recursiveSum(root, 0, &result)
	return result
}

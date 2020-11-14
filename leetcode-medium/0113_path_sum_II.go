package medium

// TreeNode : binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findAllPath(root *TreeNode, currentSet []int, currentCount int, sum int, result *[][]int) {
	if root == nil {
		if currentCount == sum {
			cp := make([]int, len(currentSet))
			copy(cp, currentSet)
			*result = append(*result, cp)
		}
		return
	}
	currentCount += root.Val
	currentSet = append(currentSet, root.Val)

	if root.Left != nil {
		findAllPath(root.Left, currentSet, currentCount, sum, result)
	}
	if root.Right != nil {
		findAllPath(root.Right, currentSet, currentCount, sum, result)
	}
	if root.Left == nil && root.Right == nil {
		findAllPath(root.Left, currentSet, currentCount, sum, result)
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	if root != nil && sum != 0 {
		findAllPath(root, []int{}, 0, sum, &result)
	}
	return result
}

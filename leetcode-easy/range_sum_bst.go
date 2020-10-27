package easy

/*
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.
*/

// func recursion(root *TreeNode, L int, R int, sum int)

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}

	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}

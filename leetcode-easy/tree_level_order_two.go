package easy

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values.
(ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
Memory Usage: 3 MB, less than 5.84% of Go online submissions for Binary Tree Level Order Traversal II.
*/

func reverseOrder(resultMap *[][]int) {
	length := len(*resultMap)
	for i := 0; i < length; i++ {
		(*resultMap)[i], (*resultMap)[length-1] = (*resultMap)[length-1], (*resultMap)[i]
		length--
	}
}

func traversalBottom(root *TreeNode, resultMap *[][]int, depth int) {
	if root == nil {
		return
	}

	if len(*resultMap) <= depth {
		*resultMap = append(*resultMap, []int{})
	}

	(*resultMap)[depth] = append((*resultMap)[depth], root.Val)

	traversalBottom(root.Left, resultMap, depth+1)
	traversalBottom(root.Right, resultMap, depth+1)
}

// LevelOrderBottom :
func LevelOrderBottom(root *TreeNode) [][]int {
	var resultMap [][]int
	traversalBottom(root, &resultMap, 0)
	reverseOrder(&resultMap)
	return resultMap
}

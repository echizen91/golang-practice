package easy

import "math"

/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Solution:
Runtime: 4 ms, faster than 93.34% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.4 MB, less than 19.13% of Go online submissions for Maximum Depth of Binary Tree.
*/

// MaxDepthTree : find the depth of a tree
func MaxDepthTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(MaxDepthTree(root.Left)), float64(MaxDepthTree(root.Right))))

}

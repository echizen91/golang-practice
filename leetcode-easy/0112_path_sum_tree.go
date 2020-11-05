package easy

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Solution:
Runtime: 4 ms, faster than 95.75% of Go online submissions for Path Sum.
Memory Usage: 4.8 MB, less than 100.00% of Go online submissions for Path Sum.
*/

func findSum(n *TreeNode, current, sum int) bool {
	res := false
	if n == nil {
		return res
	}
	current += n.Val
	if n.Left == nil && n.Right == nil {
		if current == sum {
			return true
		}
		return res
	}
	if n.Left != nil {
		res = res || findSum(n.Left, current, sum)
	}
	if n.Right != nil {
		res = res || findSum(n.Right, current, sum)
	}
	return res
}

func hasPathSum(root *TreeNode, sum int) bool {
	return findSum(root, 0, sum)
}

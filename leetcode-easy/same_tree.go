package easy

/*
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Same Tree.
*/

// TreeNode : Create Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SameTree :
func SameTree(p *TreeNode, q *TreeNode) bool {
	if p == q && p == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right)
	}

	return false
}

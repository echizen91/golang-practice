package easy

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Symmetric Tree.
*/

func checkSymmetry(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}
	return checkSymmetry(p.Left, q.Right) && checkSymmetry(p.Right, q.Left)
}

// SymmetricTree : check if tree is symmetrical
func SymmetricTree(root *TreeNode) bool {
	return root == nil || checkSymmetry(root.Left, root.Right)
}

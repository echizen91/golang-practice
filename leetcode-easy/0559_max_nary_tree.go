package easy

/*
Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Depth of N-ary Tree.
Memory Usage: 3.3 MB, less than 8.79% of Go online submissions for Maximum Depth of N-ary Tree.
*/

// Node tree
type Node struct {
	Val      int
	Children []*Node
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMax(n *Node) int {
	if n == nil {
		return 0
	}

	d := 1
	for _, c := range n.Children {
		d = max(d, findMax(c)+1)
	}

	return d
}

func maxDepth(root *Node) int {
	return findMax(root)
}

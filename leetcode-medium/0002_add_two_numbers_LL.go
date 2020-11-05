package medium

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Solution:
Runtime: 12 ms, faster than 66.70% of Go online submissions for Add Two Numbers.
Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Add Two Numbers.
*/

// ListNode : Linked List Nodes
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbersLL : Add numbers in reverse order (only works for int32 numbers)
func AddTwoNumbersLL(l1 *ListNode, l2 *ListNode) *ListNode {
	var first int
	var second int
	initial := 1
	for l1.Next != nil {
		first += initial * l1.Val
		l1 = l1.Next
		initial *= 10
	}
	first += initial * l1.Val
	initial = 1
	for l2.Next != nil {
		second += initial * l2.Val
		l2 = l2.Next
		initial *= 10
	}
	second += initial * l2.Val

	totalVal := first + second

	result := &ListNode{}
	tempNode := &ListNode{}
	x := 1
	for totalVal > 0 {
		remainder := totalVal % 10
		totalVal /= 10
		if x == 1 {
			result.Val = remainder
			x++
			continue
		}
		if x == 2 {
			tempNode.Val = remainder
			result.Next = tempNode
			x++
			continue
		}

		next := &ListNode{
			Val: remainder,
		}
		tempNode.Next = next
		tempNode = tempNode.Next
	}

	return result
}

// AddTwoNumbersLLOptimized : Add numbers in reverse order (Done)
func AddTwoNumbersLLOptimized(l1 *ListNode, l2 *ListNode) *ListNode {
	var carryOver, l1Count, l2Count int

	temp1 := *&l1
	temp2 := *&l2
	for temp1.Next != nil {
		l1Count++
		temp1 = temp1.Next
	}
	for temp2.Next != nil {
		l2Count++
		temp2 = temp2.Next
	}

	if l1Count < l2Count {
		l2, l1 = l1, l2
	}
	l1Head := *&l1
	l2Head := *&l2

	for l1Head.Next != nil && l2Head.Next != nil {
		total := l1Head.Val + l2Head.Val + carryOver
		remainder := total % 10
		carryOver = total / 10
		l1Head.Val = remainder

		// Next for both
		l1Head = l1Head.Next
		l2Head = l2Head.Next
	}

	if l1Head.Next == nil {
		goto end
	}

	{
		total := l1Head.Val + l2Head.Val + carryOver
		remainder := total % 10
		carryOver = total / 10
		l1Head.Val = remainder

		l1Head = l1Head.Next
		l2Head = l2Head.Next
	}

	for l1Head.Next != nil {
		total := l1Head.Val + carryOver
		remainder := total % 10
		carryOver = total / 10
		l1Head.Val = remainder

		// Next
		l1Head = l1Head.Next
	}
end:
	{
		var total int
		if l2Head != nil {
			total = l1Head.Val + l2Head.Val + carryOver
		} else {
			total = l1Head.Val + carryOver
		}

		remainder := total % 10
		carryOver = total / 10
		l1Head.Val = remainder
	}

	if carryOver > 0 {
		l1Head.Next = &ListNode{
			Val: 1,
		}
	}

	return l1
}

package medium

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Solution:
Not Solved, Need to consider doing addition directly from linked list
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

// AddTwoNumbersLLOptimized : Add numbers in reverse order (NOT DONE)
func AddTwoNumbersLLOptimized(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Next == nil && l2.Next == nil {
		total := l1.Val + l2.Val
		remainder := total % 10
		carryOver := total / 10
		l1.Val = remainder
		if carryOver > 0 {
			next := &ListNode{
				Val: 1,
			}
			l1.Next = next
		}
		return l1
	}
	result := &ListNode{}
	tempNode := &ListNode{}
	var carryOver int
	x := 1
	for l1.Next != nil && l2.Next != nil {
		if x == 1 {
			total := l1.Val + l2.Val
			remainder := total % 10
			carryOver = total / 10
			result.Val = remainder
			x++
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		if x == 2 {
			total := l1.Val + l2.Val + carryOver
			remainder := total % 10
			carryOver = total / 10
			tempNode.Val = remainder
			result.Next = tempNode
			x++
			l1 = l1.Next
			l2 = l2.Next
			continue
		}

		total := l1.Val + l2.Val + carryOver
		remainder := total % 10
		carryOver = total / 10
		next := &ListNode{
			Val: remainder,
		}
		tempNode.Next = next
		tempNode = tempNode.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	{
		total := l1.Val + l2.Val + carryOver
		remainder := total % 10
		carryOver = total / 10
		next := &ListNode{
			Val: remainder,
		}
		tempNode.Next = next
		tempNode = tempNode.Next
	}
	{
		// Rest
		for l1.Next != nil {
			total := l1.Val + carryOver
			remainder := total % 10
			carryOver = total / 10
			next := &ListNode{
				Val: remainder,
			}
			tempNode.Next = next
			tempNode = tempNode.Next
			l1 = l1.Next
		}
		for l2.Next != nil {
			total := l2.Val + carryOver
			remainder := total % 10
			carryOver = total / 10
			next := &ListNode{
				Val: remainder,
			}
			tempNode.Next = next
			tempNode = tempNode.Next
			l2 = l2.Next
		}
	}

	if carryOver > 0 {
		next := &ListNode{
			Val: 1,
		}
		tempNode.Next = next
	}
	return result
}

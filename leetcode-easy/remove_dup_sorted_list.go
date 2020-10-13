package easy

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Solution:
Runtime: 4 ms, faster than 90.46% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
*/

// ListNode : LinkedList
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveDupSortedList : Make sure all duplicates are deleted
func RemoveDupSortedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tempHead := *&head
	for tempHead.Next != nil {
		if tempHead.Val == tempHead.Next.Val {
			tempHead.Next = tempHead.Next.Next
			continue
		}
		tempHead = tempHead.Next
	}

	return head
}

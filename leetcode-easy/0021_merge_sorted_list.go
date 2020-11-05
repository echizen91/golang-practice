package easy

/*
Merge two sorted linked lists and return it as a new sorted list.
The new list should be made by splicing together the nodes of the first two lists.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.
*/

// MergeSortedList : merge two sorted list into one
func MergeSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = MergeSortedList(l1.Next, l2)
		return l1
	}
	l2.Next = MergeSortedList(l1, l2.Next)
	return l2

}

package easy

/*
Reverse a singly linked list.
*/

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		current, previous, current.Next = current.Next, current, previous
	}
	return previous
}

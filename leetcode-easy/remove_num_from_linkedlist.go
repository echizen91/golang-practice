package easy

/*
Remove all elements from a linked list of integers that have value val.
*/

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
Runtime: 8 ms, faster than 87.68% of Go online submissions for Remove Linked List Elements.
Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}

	if head.Val == val {
		return head.Next
	}
	return head
}

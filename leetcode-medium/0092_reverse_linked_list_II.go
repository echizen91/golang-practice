package medium

/*
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	conn := &ListNode{0, head}

	//  dummy head is used for edge case like 1->2 m=1, n=2
	dh := conn

	for m != 1 {
		conn = head
		head = head.Next
		m = m - 1
		n = n - 1
	}

	for n != 1 {
		tail := head.Next
		head.Next = head.Next.Next
		tail.Next = conn.Next
		conn.Next = tail
		n = n - 1
	}

	return dh.Next
}

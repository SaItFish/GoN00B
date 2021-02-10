package offer

type ListNode18 struct {
	Val  int
	Next *ListNode18
}

func deleteNode18(head *ListNode18, val int) *ListNode18 {
	move := head
	if head.Val == val {
		return head.Next
	}
	for move.Next != nil {
		if move.Next.Val == val {
			move.Next = move.Next.Next
			return head
		}
		move = move.Next
	}
	return head
}

package offer

type ListNode22 struct {
	Val  int
	Next *ListNode22
}

func getKthFromEnd(head *ListNode22, k int) *ListNode22 {
	if head == nil {
		return nil
	}
	var first, last = head, head
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		last = last.Next
	}
	return last
}

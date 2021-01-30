package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

// 344. 反转字符串
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {

}

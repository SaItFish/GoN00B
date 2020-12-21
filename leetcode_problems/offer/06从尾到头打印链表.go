package offer

type ListNode06 struct {
	Val  int
	Next *ListNode06
}

func reversePrint06(head *ListNode06) []int {
	if head == nil {
		return []int{}
	}
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return reverse06(res)
}

func reverse06(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < len(nums)/2 {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

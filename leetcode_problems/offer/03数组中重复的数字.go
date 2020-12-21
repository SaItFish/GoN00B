package offer

func findRepeatNumber03(nums []int) int {
	for i, x := range nums {
		if i == x {
			continue
		}
		if nums[x] == x {
			return x
		}
		nums[x], nums[i] = x, nums[x]
	}
	return 0
}

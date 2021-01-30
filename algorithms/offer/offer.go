// Package offer
// @file: offer.go
// @date: 2021/1/15
package offer

// 剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	votes := 0
	var res int
	for _, num := range nums {
		if votes == 0 {
			res = num
		}
		if num == res {
			votes++
		} else {
			votes--
		}
	}
	return res
}

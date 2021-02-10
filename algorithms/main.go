// package main is program entry
// @file: main.go
// @date: 2020/9/11
package main

import (
	"sort"
)

func main() {

}

func eat(a, b, c int) int {
	nums := []int{a, b, c}
	sort.Ints(nums)
	sum := a + b + c
	if nums[0]+nums[1] > nums[2] {
		return (sum + 2) / 3
	}
	if 2*(nums[0]+nums[1]) < nums[2] {
		return (nums[2] + 1) / 2
	}
	return (sum + 2) / 3
}

// Package labuladong
// @file: labuladong.go
// @date: 2021/2/10
package labuladong

import (
	"math"
	"sort"
)

// 969. 煎饼排序
// 猜想：如果希望求最小的翻转方法，可以通过每次反转造成逆序数变化的最大值作为参考值，结合动态规划
func pancakeSort(arr []int) []int {
	reverse := func(nums []int, i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	findMaxIdx := func(nums []int, n int) int {
		idx := 0
		for i := 0; i < n; i++ {
			if nums[i] > nums[idx] {
				idx = i
			}
		}
		return idx
	}

	res := make([]int, 0)
	var cakeSort func([]int, int)
	// sort会将前n个烧饼排序
	cakeSort = func(nums []int, n int) {
		if n == 0 {
			return
		}
		idx := findMaxIdx(nums, n)
		if idx != 0 {
			reverse(nums, 0, idx)
			res = append(res, idx+1)
		}
		reverse(nums, 0, n-1)
		res = append(res, n)
		cakeSort(nums, n-1)
	}

	cakeSort(arr, len(arr))
	return res
}

func eatGrape(a, b, c int) int {
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

// 43. 字符串相乘
func multiply(num1 string, num2 string) string {
	var zero byte = '0'
	if num1[0] == zero || num2[0] == zero {
		return "0"
	}
	res := make([]byte, len(num1)+len(num2))

	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {
			root := i + j + 1
			product := (num1[j]-zero)*(num2[i]-zero) + res[root] // 这里加上个位值
			a, b := product%10, product/10
			res[root] = a
			res[root-1] += b
		}
	}

	for i := 0; i < len(res); i++ {
		res[i] += zero
	}

	if res[0] == zero {
		return string(res[1:])
	}
	return string(res)
}

// 224. 基本计算器
func calculate(s string) int {
	isDigit := func(c byte) bool {
		if '0' <= c && c <= '9' {
			return true
		}
		return false
	}

	sum := func(nums []int) int {
		res := 0
		for i := 0; i < len(nums); i++ {
			res += nums[i]
		}
		return res
	}

	n := len(s)
	i := 0
	var helper func(string) int
	helper = func(s string) int {
		var num = 0
		var sign byte = '+'
		stack := make([]int, 0)
		for i < n {
			c := s[i]
			i++
			if isDigit(c) {
				num = num*10 + int(c-'0')
			}
			// 左括号递归计算括号内内容
			if c == '(' {
				num = helper(s)
			}
			// 不为数字，不为空  只可能是符号或者到达字符串末尾
			if (!isDigit(c) && c != ' ') || i == n {
				switch sign {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				}
				sign = c
				num = 0
			}
			// 右括号直接返回结果
			if c == ')' {
				break
			}
		}
		return sum(stack)
	}

	return helper(s)
}

// 227. 基本计算器 II
func calculate2(s string) int {
	isDigit := func(c byte) bool {
		if '0' <= c && c <= '9' {
			return true
		}
		return false
	}

	sum := func(nums []int) int {
		res := 0
		for i := 0; i < len(nums); i++ {
			res += nums[i]
		}
		return res
	}

	n := len(s)
	num := 0
	var sign byte = '+'
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		c := s[i]
		if isDigit(c) {
			num = num*10 + int(c-'0')
		}
		if (!isDigit(c) && c != ' ') || i == n-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pre := stack[len(stack)-1]
				stack[len(stack)-1] = pre * num
			case '/':
				pre := stack[len(stack)-1]
				stack[len(stack)-1] = pre / num
			}
			sign = c
			num = 0
		}

	}

	return sum(stack)
}

// 42. 接雨水
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	n := len(height)
	left, right := 0, n-1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	lMax, rMax := height[left], height[right]
	for left <= right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])
		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}
	return res
}

// 20. 有效的括号
func isValidBrackets(s string) bool {
	stack := make([]byte, 0)
	corresponding := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if corresponding[s[i]] != pre {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// 391. 完美矩形
func isRectangleCover(rectangles [][]int) bool {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	X1, Y1 := math.MaxInt64, math.MaxInt64
	X2, Y2 := math.MinInt64, math.MinInt64

	actualArea := 0
	points := make(map[[2]int]struct{}, 0)

	for _, v := range rectangles {
		x1, y1, x2, y2 := v[0], v[1], v[2], v[3]
		X1, Y1 = min(X1, x1), min(Y1, y1)
		X2, Y2 = max(X2, x2), max(Y2, y2)

		actualArea += (x2 - x1) * (y2 - y1)
		ps := [][2]int{{x1, y1}, {x1, y2}, {x2, y1}, {x2, y2}}
		for _, p := range ps {
			if _, ok := points[p]; ok {
				delete(points, p)
			} else {
				points[p] = struct{}{}
			}
		}
	}
	expectedArea := (X2 - X1) * (Y2 - Y1)
	// 面积
	if expectedArea != actualArea {
		return false
	}
	// 顶点
	if len(points) != 4 {
		return false
	}
	// 最终的四个顶点必须是完美矩形的理论坐标
	ps := [][2]int{{X1, Y1}, {X1, Y2}, {X2, Y1}, {X2, Y2}}
	for _, p := range ps {
		if _, ok := points[p]; !ok {
			return false
		}
	}
	return true
}

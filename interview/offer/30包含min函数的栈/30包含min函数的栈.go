// Package min_stack
// @file: 30包含min函数的栈.go
// @date: 2021/2/16
package min_stack

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.min) != 0 && s.min[len(s.min)-1] > x {
		s.min = append(s.min, x)
	}
}

func (s *MinStack) Pop() {
	t := s.Top()
	s.stack = s.stack[:len(s.stack)-1]
	if t == s.Min() {
		s.min = s.min[:len(s.min)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Min() int {
	return s.min[len(s.min)-1]
}

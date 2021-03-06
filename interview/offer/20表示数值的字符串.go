// Package offer
// @file: 20表示数值的字符串.go
// @date: 2021/2/14
package offer

/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

有限状态自动机
*/

func isNumberX20(s string) bool {
	status := []map[byte]int{
		// s: 符号, d: 数字
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0 开始的空格
		{'d': 2, '.': 4},                 // 1 幂符号前的正负号
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2 小数点前的数字 √
		{'d': 3, 'e': 5, ' ': 8},         // 3 小数点、小数点后的数字 √
		{'d': 3},                         // 4 当小数点前为空格时，小数点、小数点后的数字
		{'s': 6, 'd': 7},                 // 5 幂符号
		{'d': 7},                         // 6 幂符号后的正负号
		{'d': 7, ' ': 8},                 // 7 幂符号后的数字 √
		{' ': 8},                         // 8 结尾的空格 √
	}

	p := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		var t byte
		switch {
		case '0' <= c && c <= '9':
			t = 'd'
		case c == '+' || c == '-':
			t = 's'
		case c == 'e' || c == 'E':
			t = 'e'
		case c == '.' || c == ' ':
			t = c
		default:
			t = '?'
		}
		if _, ok := status[p][t]; !ok {
			return false
		}
		p = status[p][t]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}

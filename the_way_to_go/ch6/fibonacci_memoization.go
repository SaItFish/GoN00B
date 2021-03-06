// Package ch6 is chapter 6
// @file: fibonacci_memoization.go
// @description: 使用内存存储数据，优化程序运行时间
// @author: SaltFish
// @date: 2020/08/06
package ch6

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

// FibonacciMem is fun
func FibonacciMem() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

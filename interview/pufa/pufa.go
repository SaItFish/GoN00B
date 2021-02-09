// Package pufa
// @file: pufa.go
// @date: 2021/2/8
package pufa

// 判断是否为质数
func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 计算n以内的质数个数
func countPrimes(n int) int {
	primes := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if primes[i] == false {
			for j := i * i; j < n; j += i {
				primes[j] = true
			}
			count++
		}
	}
	return count
}

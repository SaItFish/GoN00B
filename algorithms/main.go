// package main is program entry
// @file: main.go
// @date: 2020/9/11
package main

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := make([]int, 0)
L:
	for i := 2; i <= n; i++ {
		for _, j := range primes {
			if i%j == 0 {
				continue L
			}
		}
		primes = append(primes, i)
	}
	return len(primes)
}

func main() {

}

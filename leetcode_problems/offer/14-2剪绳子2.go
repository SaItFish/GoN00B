package offer

func cuttingRope14(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b, p, x, rem := n/3-1, n%3, 1000000007, 3, 1
	for a > 0 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = x * x % p
		a /= 2
	}
	if b == 0 {
		return rem * 3 % p
	}
	if b == 1 {
		return rem * 4 % p
	}
	return rem * 6 % p
}

// 求 (x^a) % p —— 循环求余法
func remainder14(x, a, p int) int {
	rem := 1
	for i := 0; i < a; i++ {
		rem = (rem * x) % p
	}
	return rem
}

// 快速幂求余
func fastRemainder14(x, a, p int) int {
	rem := 1
	for a > 0 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = x * x % p
		a /= 2
	}
	return rem
}

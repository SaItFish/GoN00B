package offer

func myPow16(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	var res = 1.0
	for n != 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

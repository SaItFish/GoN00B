package offer

import "math"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	switch a, b := n/3, n%3; b {
	case 0:
		return int(math.Pow(3, float64(a)))
	case 1:
		return int(math.Pow(3, float64(a-1)) * 4)
	default:
		return int(math.Pow(3, float64(a)) * 2)
	}
}

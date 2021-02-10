package offer

func hammingWeight15(num uint32) int {
	res := 0
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

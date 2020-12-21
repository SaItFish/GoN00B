package offer

func printNumbers17(n int) []int {
	cnt := 1
	for i := 0; i < n; i++ {
		cnt *= 10
	}
	res := make([]int, cnt-1)
	for i := range res {
		res[i] = i + 1
	}
	return res
}

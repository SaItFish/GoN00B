package offer

func movingCount13(m int, n int, k int) int {
	var (
		mark  = make([][]int, m)
		d     = []int{-1, 0, 1, 0, -1}
		r     = 0
		check func(int) int
		dfs   func(int, int)
	)
	// 初始化二维数组
	for i := range mark {
		mark[i] = make([]int, n)
	}

	// 计算数位和
	check = func(num int) int {
		res := 0
		for num != 0 {
			res += num % 10
			num /= 10
		}
		return res
	}

	dfs = func(x int, y int) {
		// 是否越界或已经走过
		if x < 0 || x >= m || y < 0 || y >= n || mark[x][y] == 1 {
			return
		}
		// 题目要求
		if check(x)+check(y) > k {
			return
		}
		// 标记走过的位置
		mark[x][y] = 1
		// 计数
		r++

		for i := 0; i < 4; i++ {
			dfs(x+d[i], y+d[i+1])
		}
	}
	dfs(0, 0)
	return r
}

package offer

func exist12(board [][]byte, word string) bool {
	var (
		dfs    func(int, int, int) bool
		d      = []int{-1, 0, 1, 0, -1}
		rows   = len(board)
		cols   = len(board[0])
		length = len(word) - 1
	)
	dfs = func(x, y, pos int) bool {
		// 是否越界
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return false
		}
		// 保存当前字符
		ch := board[x][y]
		// 字符是否匹配
		if ch == '#' || ch != word[pos] {
			return false
		}
		// 完全匹配word
		if pos == length {
			return true
		}
		// 标记当前走过的字符
		board[x][y] = '#'
		for i := 0; i < 4; i++ {
			if dfs(x+d[i], y+d[i+1], pos+1) {
				return true
			}
		}
		board[x][y] = ch
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

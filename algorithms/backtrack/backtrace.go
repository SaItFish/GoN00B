// Package backtrack
// @file: backtrace.go
// @date: 2021/1/15
package backtrack

/*
回溯法
result = []
func backtrack(选择列表,路径):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(选择列表,路径)
        撤销选择
*/
// 78. 子集
func subsets(nums []int) [][]int {
	length := len(nums)
	res := make([][]int, 0)

	var backtrack func(int, []int)
	backtrack = func(pos int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := pos; i < length; i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{})
	return res
}

// 46. 全排列
func permute(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}

	var dfs func([]int)
	dfs = func(path []int) {
		// 结束条件
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for _, num := range nums {
			if visited[num] {
				continue
			}
			// 做选择
			path = append(path, num)
			visited[num] = true
			// 下一个决策
			dfs(path)
			// 撤销选择
			visited[num] = false
			path = path[:len(path)-1]
		}
	}

	dfs([]int{})
	return res
}

// 51. N 皇后
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	isValid := func(board [][]byte, row, col int) bool {
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' { // 上方
				return false
			}
			if col-row+i >= 0 && board[i][col-row+i] == 'Q' { //左上方
				return false
			}
			if col+row-i < n && board[i][col+row-i] == 'Q' { //右上方
				return false
			}
		}
		return true
	}

	var backTrace func([][]byte, int)
	backTrace = func(board [][]byte, row int) {
		if row == len(board) {
			t := make([]string, row)
			for k, bs := range board {
				t[k] = string(bs)
			}
			result = append(result, t)
			return
		}
		for col := 0; col < len(board[row]); col++ {
			if !isValid(board, row, col) {
				continue
			}
			// 做选择
			board[row][col] = 'Q'
			// 进入下一次决策
			backTrace(board, row+1)
			// 撤销选择
			board[row][col] = '.'
		}
	}

	backTrace(board, 0)
	return result
}

func PrintRes() {
	solveNQueens(4)
}

package offer

func findNumberIn2DArray04(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		switch {
		case matrix[i][j] == target:
			return true
		case matrix[i][j] < target:
			i++
		case matrix[i][j] > target:
			j--
		}
	}
	return false
}

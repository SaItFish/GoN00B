package offer

func isMatch19(s string, p string) bool {
	s, p = "#"+s, "#"+p
	m, n := len(s), len(p)
	var dp = make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == 0 {
				dp[i][j] = j > 1 && p[j] == '*' && dp[i][j-2]
			} else if p[j] == s[i] || p[j] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j] == '*' {
				dp[i][j] = j > 1 && dp[i][j-2] || (p[j-1] == s[i] || p[j-1] == '.') && dp[i-1][j]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[m-1][n-1]
}

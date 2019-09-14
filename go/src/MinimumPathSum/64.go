package MinimumPathSum

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				if j > 0 {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				} else {
					dp[i][j] = grid[i][j]
				}
			} else if j == 0 {
				if i > 0 {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				} else {
					dp[i][j] = grid[i][j]
				}
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				} else {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				}
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

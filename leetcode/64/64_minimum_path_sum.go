package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	dp := [][]int{}
	m := len(grid)
	n := len(grid[0])
	// It can be solved without allocation of new array, how?
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

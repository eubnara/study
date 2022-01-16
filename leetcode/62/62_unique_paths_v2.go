package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := [][]int{}
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}
	dp[0][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
func main() {
	fmt.Println(uniquePaths(3, 7))
	// fmt.Println(uniquePaths(23, 12))
}

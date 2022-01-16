package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := [][]int{}
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}
	dp[0][1] = 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {

			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0},
	}))
}

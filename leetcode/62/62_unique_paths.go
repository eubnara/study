package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := [][]int{}
	visited := [][]bool{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
		visited = append(visited, make([]bool, n))
	}
	dp[0][0] = 1
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		y := queue[0][0]
		x := queue[0][1]
		queue = queue[1:]
		if visited[y][x] {
			continue
		}
		visited[y][x] = true
		if x+1 < n {
			dp[y][x+1] += dp[y][x]
			queue = append(queue, [2]int{y, x + 1})
		}
		if y+1 < m {
			dp[y+1][x] += dp[y][x]
			queue = append(queue, [2]int{y + 1, x})
		}
	}
	return dp[m-1][n-1]
}
func main() {
	fmt.Println(uniquePaths(3, 7))
	// fmt.Println(uniquePaths(23, 12))
}

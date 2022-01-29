package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i:=2;i<=n;i++ {
		for j:=1;j<=i;j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3)) // 5
	fmt.Println(numTrees(19));
}
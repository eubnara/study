package main

import "fmt"

func numTrees(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
    return numTreesHelper(1, n, dp)
}

func numTreesHelper(start, end int, dp [][]int) int {
	if start == end {
		return 1
	}
	if dp[start][end] != 0 {
		return dp[start][end]
	}
	ans := 0
	for i:=start;i<=end;i++ {
		left := 1
		right := 1
		if start+1 < i {
			left = numTreesHelper(start, i-1, dp)
		}
		if i < end-1 {
			right = numTreesHelper(i+1, end, dp)
		}
		ans += left * right
	}
	dp[start][end] = ans
	return ans
}

func main() {
	fmt.Println(numTrees(3)) // 5
	fmt.Println(numTrees(19));
}
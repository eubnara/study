package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistanceHelper(word1, word2 string, ai, bi int, dp *[501][501]int) int {
	al := len(word1)
	bl := len(word2)
	if dp[ai][bi] != 0 {
		return dp[ai][bi]
	}
	ret := 501
	if al == ai {
		return bl - bi
	}
	if bl == bi {
		return al - ai
	}
	if word1[ai] == word2[bi] {
		ret = min(ret, minDistanceHelper(word1, word2, ai+1, bi+1, dp)) // pass
	} else {
		ret = min(ret, minDistanceHelper(word1, word2, ai+1, bi+1, dp)+1) // replace
	}
	ret = min(ret, minDistanceHelper(word1, word2, ai+1, bi, dp)+1) // delete
	ret = min(ret, minDistanceHelper(word1, word2, ai, bi+1, dp)+1) // insert
	dp[ai][bi] = ret
	return ret
}

func minDistance(word1 string, word2 string) int {
	dp := [501][501]int{}
	if word1 == word2 {
		return 0
	}
	return minDistanceHelper(word1, word2, 0, 0, &dp)
}
func main() {
	// fmt.Println(minDistance("horse", "ros")) // 3
	// fmt.Println(minDistance("intention", "execution")) // 5
	fmt.Println(minDistance("sea", "eat")) // 2
}
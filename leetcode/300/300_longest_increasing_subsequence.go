package main

import (
	"fmt"
	"os"
	"strconv"
)

func binarySearch(dp []int, n int) int {
	l := len(dp)
	if l == 0 {
		return 0
	}
	if dp[l/2] < n {
		return l/2 + binarySearch(dp[l/2+1:], n) + 1
	} else if dp[l/2] == n {
		if l == 1 {
			return 0
		}
		return l/2 + binarySearch(dp[l/2:], n)
	} else {
		return binarySearch(dp[:l/2], n)
	}
}

func lengthOfLIS(nums []int) int {
	dp := []int{-10001}
	for _, v := range nums {
		pos := binarySearch(dp, v)
		l := len(dp)
		if pos >= l {
			dp = append(dp, v)
		} else if dp[pos] > v {
			dp[pos] = v
		}
	}
	return len(dp) - 1
}

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	args := []int{}
	for _, v := range os.Args[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		args = append(args, num)
	}
	fmt.Println(lengthOfLIS(args))
}

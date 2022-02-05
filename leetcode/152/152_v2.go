package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// https://leetcode.com/problems/maximum-product-subarray/discuss/550573/go-O(N)
func maxProduct(nums []int) int {
	ret, curMin, curMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v < 0 {
			curMin, curMax = curMax, curMin
		}
		curMin = min(v, curMin*v)
		curMax = max(v, curMax*v)
		if curMax > ret {
			ret = curMax
		}
	}

	return ret
}

func main() {
	// fmt.Println(maxProduct([]int{-2})) // -2
	// fmt.Println(maxProduct([]int{-3, -1, -1})) // 3
	fmt.Println(maxProduct([]int{0, 2})) // 2
}

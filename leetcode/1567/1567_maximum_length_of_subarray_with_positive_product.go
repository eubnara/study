package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func getMaxLen(nums []int) int {
	ans, nIdx, pIdx := 0, -1, -1
	for i, v := range nums {
		if v == 0 {
			nIdx, pIdx = -1, -1
			continue
		}
		if v < 0 {
			nIdx, pIdx = pIdx, nIdx
			if nIdx == -1 {
				nIdx = i
			}
		} else {
			if pIdx == -1 {
				pIdx = i
			}
		}
		if pIdx != -1 {
			ans = max(ans, i-pIdx+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(getMaxLen([]int{0, 1, -2, -3, -4})) // 3
}

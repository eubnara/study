package main

import (
	"fmt"
	"sort"
)

func max (a, b int) int {
    if a < b {
        return b
    }
    return a
}
func deleteAndEarn(nums []int) int {
    dp := map[int]int{}
    sort.Sort(sort.IntSlice(nums))
    return deleteAndEarnHelper(&nums, 0, 0, &dp)
}
func deleteAndEarnHelper(nums *[]int, curIdx, curMin int, dp *map[int]int) int {
    if v, exists := (*dp)[curMin];exists {
        return v
    }
    ans := 0
    l := len(*nums)
	idx := curIdx
	for idx < l {
		if curMin <= (*nums)[idx] {
			break
		}
		idx++
	}
	if idx >= l {
		return ans
	}
	curMin = (*nums)[idx]
	accum := 0
	for idx < l {
		if (*nums)[idx] != curMin {
			break
		}
		accum += curMin
		idx++
	}
	ans = max(accum + deleteAndEarnHelper(nums, idx, curMin+2, dp), deleteAndEarnHelper(nums, idx, curMin, dp))
	(*dp)[curMin] = ans
	return ans
}

func main() {
	fmt.Println(deleteAndEarn([]int{3,4,2}))
}
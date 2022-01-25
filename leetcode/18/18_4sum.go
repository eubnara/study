package leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return kSum(nums, target, 4)
}

func kSum(nums []int, target int, k int) [][]int {
	ans := [][]int{}
	m := map[int]bool{}
	if k != 2 {
		for i, v := range nums {
			if m[v] {
				continue
			}
			m[v] = true
			subAns := kSum(nums[i+1:], target-v, k-1)
			for _, s := range subAns {
				ans = append(ans, append(s, v))
			}
		}
	} else {
		found := map[int]bool{}
		for i := 0; i < len(nums); i++ {
			v := nums[i]
			complement := target - v
			if found[complement] {
				ans = append(ans, []int{v, complement})
				for i+1 < len(nums) {
					if v != nums[i+1] {
						break
					}
					i++
				}
			}
			found[v] = true
		}
	}
	return ans
}

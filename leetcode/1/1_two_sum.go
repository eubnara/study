package leetcode

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	sortedNums := []int{}
	for _, v := range nums {
		sortedNums = append(sortedNums, v)
	}
	sort.Slice(sortedNums, func(a, b int) bool {
		return sortedNums[a] < sortedNums[b]
	})
	i, j := 0, len(sortedNums)-1
	ans := []int{}
	for i < j {
		sum := sortedNums[i] + sortedNums[j]
		if sum == target {
			for idx := 0; idx < len(nums); idx++ {
				if nums[idx] == sortedNums[i] {
					ans = append(ans, idx)
					break
				}
			}
			for idx := len(nums) - 1; idx >= 0; idx-- {
				if nums[idx] == sortedNums[j] {
					ans = append(ans, idx)
					break
				}
			}
			return ans
		}
		if sum < target {
			i++
			continue
		}
		if sum > target {
			j--
			continue
		}
	}
	return ans
}

package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	visited := map[int]bool{}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		if visited[v] {
			continue
		}
		visited[v] = true
		subAns := twoSum(nums[:i], -v)
		for _, s := range subAns {
			ans = append(ans, append(s, v))
		}
	}
	return ans
}

func twoSum(nums []int, target int) [][]int {
	// nums is sorted.
	ret := [][]int{}
	p, q := 0, len(nums)-1
	for p < q {
		lval, rval := nums[p], nums[q]
		sum := lval + rval
		if sum == target {
			ret = append(ret, []int{lval, rval})
			p++
			for p < len(nums) {
				if nums[p] != lval {
					break
				}
				p++
			}
			q--
			for q > 0 {
				if nums[q] != rval {
					break
				}
				q--
			}
			continue
		}
		if sum < target {
			p++
			for p < len(nums) {
				if nums[p] != lval {
					break
				}
				p++
			}
		} else {
			q--
			for q > 0 {
				if nums[q] != rval {
					break
				}
				q--
			}
		}
	}
	return ret
}

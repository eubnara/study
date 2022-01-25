package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return kSum(nums, target, 0, 4)
}

func kSum(nums []int, target int, start int, k int) [][]int {
	ans := [][]int{}
	m := map[int]bool{}
	if k != 2 {
		for i := start; i < len(nums); i++ {
			v := nums[i]
			if m[v] {
				continue
			}
			if v > target/k {
				break
			}
			if nums[len(nums)-1] < target/k {
				break
			}
			m[v] = true
			subAns := kSum(nums, target-v, i+1, k-1)
			for _, s := range subAns {
				ans = append(ans, append(s, v))
			}
		}
	} else {
		lo := start
		hi := len(nums) - 1
		for lo < hi {
			lval, rval := nums[lo], nums[hi]
			sum := lval + rval
			if lval > target/2 {
				break
			}
			if rval < target/2 {
				break
			}
			if sum < target {
				lo++
				for lo < hi {
					if nums[lo] != lval {
						break
					}
					lo++
				}
			} else if sum > target {
				hi--
				for lo < hi {
					if nums[hi] != rval {
						break
					}
					hi--
				}
			} else {
				ans = append(ans, []int{lval, rval})

				lo++
				for lo < hi {
					if nums[lo] != lval {
						break
					}
					lo++
				}

				hi--
				for lo < hi {
					if nums[hi] != rval {
						break
					}
					hi--
				}
			}
		}
	}
	return ans
}

func main() {
	// fmt.Println(fourSum([]int{
	// 	-3, -2, -1, 0, 0, 1, 2, 3,
	// }, 0))

	fmt.Println(fourSum([]int{
		1, 0, -1, 0, -2, 2,
	}, 0))
}

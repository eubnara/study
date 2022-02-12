package main

import "fmt"

// brute force
func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	ans := nums[0]
    for i:=0;i<l;i++ {
		sum := 0
		if nums[i] == 0 {
			continue
		}
		if i != l-1 && nums[(i-1+l)%l] > 0 {
			continue
		}
		for j:=i;j<i+l;j++ {
			idx := j % l
			sum += nums[idx]
			if ans < sum {
				ans = sum
			}
			if sum < 0 {
				break
			}
		}
	}
	return ans
}

func main() {
	// fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	// fmt.Println(maxSubarraySumCircular([]int{9, -4, -7, 9})) // 18
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
}
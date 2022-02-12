package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	ans := nums[0]
    for i:=0;i<l;i++ {
		sum := 0
		for j:=i;j<i+l;j++ {
			idx := j % l
			sum += nums[idx]
			if ans < sum {
				ans = sum
			}
		}
	}
	return ans
}

func main() {
	// fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{9, -4, -7, 9})) // 18
}
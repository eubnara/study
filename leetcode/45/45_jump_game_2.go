package main

import "fmt"

func jump(nums []int) int {
	l := len(nums)
	step := 0
	maxIdx := 0
	for i:=0;i<l-1;i++ {
		curMaxIdx := maxIdx
		for j:=i;j<=maxIdx;j++ {
			v := j + nums[j]
			if curMaxIdx < v {
				curMaxIdx = v
			}
		}
		step++
		i = maxIdx
		maxIdx = curMaxIdx
		if maxIdx >= l-1 {
			return step
		}
	}
	return step
}

func main() {
	// fmt.Println(jump([]int{7,0,9,6,9,6,1,7,9,0,1,2,9,0,3}))
	fmt.Println(jump([]int{1,2,3})) // 2
}
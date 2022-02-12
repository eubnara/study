package main

import "fmt"

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func rob(nums []int) int {
    cache := []int{
        0, nums[0],
    }
    l := len(nums)
    if l == 1 {
        return nums[0]
    }
	cache[0], cache[1] = cache[1], cache[0] + nums[1]
    for i:=2;i<l;i++ {
        cache[0], cache[1] = max(cache[0], cache[1]), cache[0] + nums[i]
    }
    return max(cache[0], cache[1])
}

func main() {
	fmt.Println(rob([]int{1,2,3,1}))
}
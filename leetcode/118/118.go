package main

import "fmt"

func generate(numRows int) [][]int {
	ans := [][]int{{1}}
	if numRows >= 2 {
		ans = append(ans, []int{1, 1})
	}
	for i := 3; i <= numRows; i++ {
		subAns := []int{1}
		for j := 1; j < i-1; j++ {
			subAns = append(subAns, ans[i-2][j-1]+ans[i-2][j])
		}
		subAns = append(subAns, 1)
		ans = append(ans, subAns)
	}
	return ans
}

func main() {
	fmt.Println(generate(5))
}

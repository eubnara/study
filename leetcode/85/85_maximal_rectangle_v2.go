package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalRectangle(matrix [][]byte) int {
	ret := 0
	r := len(matrix)
	c := len(matrix[0])
	dp := make([][]int, r)
	for i := range dp {
		dp[i] = make([]int, c)
	}
	for width := 1; width <= c; width++ {
		for i := width - 1; i < c; i++ {
			height := 0
			for j := 0; j < r; j++ {
				if matrix[j][i] == '0' {
					height = 0
					dp[j][i] = 0
					continue
				}
				dp[j][i] = 1
				if i > 0 {
					dp[j][i] += dp[j][i-1]
				}
				if dp[j][i] >= width {
					height += 1
					ret = max(ret, width*height)
				} else {
					height = 0
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	})) // 6
	fmt.Println(maximalRectangle([][]byte{
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'1', '1', '1'},
	})) // 3
	fmt.Println(maximalRectangle([][]byte{
		{'0', '1'},
	})) // 1
	fmt.Println(maximalRectangle([][]byte{
		{'0', '1'},
		{'0', '1'},
	})) // 2
	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0', '1', '1', '1', '0'},
		{'1', '1', '1', '0', '0', '0', '0', '0', '1'},
		{'0', '0', '1', '1', '0', '0', '0', '1', '1'},
		{'0', '1', '1', '0', '0', '1', '0', '0', '1'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '1', '1', '0', '1', '0', '0', '0', '1'},
		{'0', '1', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '0', '0', '1', '1', '1', '0', '0', '0'},
	}))
}

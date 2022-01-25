package leetcode

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
		for i := 0; i < c; i++ {
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

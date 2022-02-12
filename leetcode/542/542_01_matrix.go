package main

import "fmt"

func updateMatrix(mat [][]int) [][]int {
    m, n := len(mat), len(mat[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            visit(&mat, &visited, i, j, m, n)
        }
    }
    return mat
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func visit(mat *[][]int, visited *[][]bool, r, c, m, n int) int {
    if (*visited)[r][c] {
        return (*mat)[r][c]
    }
    ret := (*mat)[r][c]
    (*visited)[r][c] = true
    if ret == 0 {
        return ret
    }
    minVal := 1000
    if r > 0 {
        minVal = min(minVal, visit(mat, visited, r-1, c, m, n))
    }
    if r+1 < m {
        minVal = min(minVal, visit(mat, visited, r+1, c, m, n))
    }
    if c > 0 {
        minVal = min(minVal, visit(mat, visited, r, c-1, m, n))
    }
    if c+1 < n {
        minVal = min(minVal, visit(mat, visited, r, c+1, m, n))
    }
    ret = minVal + 1
	(*mat)[r][c] = ret
    return ret
}

func main() {
	// fmt.Println(updateMatrix([][]int{
	// 	{0,0,0},
	// 	{0,1,0},
	// 	{1,1,1},
	// }))
	fmt.Println(updateMatrix([][]int{
		{1,0,1,1,0,0,1,0,0,1},
		{0,1,1,0,1,0,1,0,1,1},
		{0,0,1,0,1,0,0,1,0,0},
		{1,0,1,0,1,1,1,1,1,1},
		{0,1,0,1,1,0,0,0,0,1},
		{0,0,1,0,1,1,1,0,1,0},
		{0,1,0,1,0,1,0,0,1,1},
		{1,0,0,0,1,1,1,1,0,1},
		{1,1,1,1,1,1,1,0,1,0},
		{1,1,1,1,0,1,0,0,1,1},
	}))
}
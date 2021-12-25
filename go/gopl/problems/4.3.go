package main

import "fmt"

func reverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	reverse(&a)
	fmt.Println(a)
}

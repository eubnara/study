package main

import (
	"fmt"
	"os"
	"strconv"
)

func rotate(a []int, n int) []int {
	l := len(a)
	n = n % l
	a = append(a[l-n:l], a[:l-n]...)
	return a
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	n := 2
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println(rotate(a, n))
}

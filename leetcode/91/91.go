package main

import "fmt"

func numDecodings(s string) int {
	l := len(s)
	if s[0] == '0' {
		return 0
	}
	cur, n, nn := 1, 0, 0
	for i:=0;i<l;i++ {
		c := s[i]
		if cur == 0 && n == 0 {
			break
		}
		if c == '0' {
			cur, n, nn = n, nn, 0
			continue
		}
		val := c - '0'
		if val <= 2 && i+1 < l {
			nextVal := s[i+1] - '0'
			if val == 1 || nextVal <= 6 {
				nn += cur
			}
		}
		n += cur
		cur, n, nn = n, nn, 0
	}
	return cur
}

func main() {
	fmt.Println(numDecodings("2611055971756562"))
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}

// func rotate(b []byte, n int) {
// 	l := len(b)
// 	n = n % l
// 	reverse(b[:l-n])
// 	reverse(b[l-n:])
// 	reverse(b)
// }

// func reverseUTF8(b []byte) {
// 	l := len(b)
// 	_, s1 := utf8.DecodeRune(b)
// 	_, s2 := utf8.DecodeLastRune(b)
// 	if s1 == 0 || s2 == 0 || s1 > l-s2 {
// 		return
// 	}
// 	var max int
// 	if s1 > s2 {
// 		max = s1
// 	} else {
// 		max = s2
// 	}
// 	if s1 == s2 {
// 		for idx := 0; idx < max; idx++ {
// 			b[idx], b[l-max+idx] = b[l-max+idx], b[idx]
// 		}
// 		reverseUTF8(b[max : l-max])
// 	} else if s1 > s2 {
// 		rotate(b[s1:], s2)
// 		rotate(b, l-s1)
// 		reverseUTF8(b[s2 : l-s1])
// 	} else {
// 		rotate(b[:l-s2], s1)
// 		rotate(b, s2)
// 		reverseUTF8(b[s2 : l-s1])
// 	}
// }

func main() {
	b := []byte("안녕하세요, hello, world!")
	fmt.Println(string(b))
	reverseUTF8(b)
	fmt.Println(string(b))
}

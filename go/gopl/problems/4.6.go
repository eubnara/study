package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func mergeSpace(b []byte) []byte {
	var buf bytes.Buffer

	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			next, _ := utf8.DecodeRune(b[i+size:])
			if !unicode.IsSpace(next) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}
		i += size
	}
	return buf.Bytes()
}

func printByteArr(b []byte) {
	for _, v := range b {
		fmt.Printf("%08b ", v)
	}
	fmt.Println()
}

func main() {
	origin := []byte("하         이 ")
	printByteArr(origin)
	printByteArr(mergeSpace(origin))
}

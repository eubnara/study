package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	nonLetterCounts, letterCounts := 0, 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts++
		} else {
			nonLetterCounts++
		}
	}
	fmt.Printf("letter\tnon-letter\n")
	fmt.Printf("%d\t%d\n", letterCounts, nonLetterCounts)
}

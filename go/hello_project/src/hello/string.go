package main

import "fmt"
import "unicode/utf8"

func main() {
	var a string = "안녕하세요"
	fmt.Println(a)
	fmt.Printf("%c\n", a[0])

	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))
}

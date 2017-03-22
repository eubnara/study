package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    fmt.Println(a)

    b := a[:]
    fmt.Println(b)

    b[3] = 5

    fmt.Println(a)
    fmt.Println(b)
}

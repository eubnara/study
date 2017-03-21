package main

import "fmt"

func main() {
    var a string = "initial"
    fmt.Println(a)

    var b, c = 1, 2
    fmt.Println(b, c)

    var d = true 
    fmt.Println(d)

    var e complex64 = 1 + 3i
    fmt.Println(e)

    f := "short"
    fmt.Println(f)

    var g bool
    fmt.Println(g)

    var h byte
    fmt.Println(h)

    var i uint8
    fmt.Println(i)

    fmt.Println(h == byte(i))
}

package main

import "fmt"

func main() {
    
    i := 1
    for i<= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 8; j<= 9 ; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
    fmt.Println("add 1 to ten")
    var sum int //no need to intialize, it intializes into zero
    var sum2 = 0
    var index int
    for index = 1; index  <= 100; index++ {
        sum += index
    }
    fmt.Println(sum)
    for index := 1; index  <= 100; index++ {
        sum2 += index
    }
    fmt.Println(sum2)
}

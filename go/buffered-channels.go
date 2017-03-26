package main

import "fmt"
import "time"

var c = make(chan int, 5)

func main() {
    go worker(1)
    for i := 0; i < 10; i++ {
        c <- i
        fmt.Println(i)
    }
}

func worker(id int) {
    for {
        _ = <-c
        time.Sleep(time.Second)
    }
}

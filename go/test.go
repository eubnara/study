package main

import "fmt"

func main() {
    messages := make(chan string, 1)
    go func() {
        messages <- "helloworld"
    }()
    fmt.Println(<-messages)
    go func() { fmt.Println(<-messages)}()
}

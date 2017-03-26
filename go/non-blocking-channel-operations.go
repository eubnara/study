package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    go func() {messages <- "first"}()

    select {
    case msg := <-messages:
        fmt.Println("received messages!:", msg)
//    default:
//        fmt.Println("no message received")
    }


    msg := "hello"
    select {
    case messages <- msg:
        //fmt.Println("sent messages", <-messages)
        fmt.Println("sent messages")
    default:
        fmt.Println("no message sent")
    }


    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

package main
import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "ping1" }()
    go func() { messages <- "ping2" }()
    go func() { messages <- "ping3" }()
    go func() { messages <- "ping4" }()
    go func() { messages <- "ping5" }()

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    //fmt.Println(<-messages)
    // cannot guarantee which one comes first.
    msg:= <-messages
    fmt.Println(msg)
}

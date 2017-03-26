package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{"Bob", 20})

    fmt.Println(person{ age: 30, name: "Alice"})

    fmt.Println(person{name: "Fred"})
    fmt.Println(person{age: 23})

    fmt.Println(&person{name: "Ann", age: 40})

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    fmt.Println(sp)
    fmt.Println(s)
}

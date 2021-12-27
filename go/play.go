package main

import "fmt"

func main() {
	var graph = make(map[string]map[string]bool)
	fmt.Println(graph["여기서"]["저기로"])
	m := map[string]bool{}
	fmt.Printf("%v\n", m["hi"])
}

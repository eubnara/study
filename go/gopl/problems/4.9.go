package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	// scanner := bufio.NewScanner(strings.NewReader(input))
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))

	scanner.Split(bufio.ScanWords)
	m := map[string]int{}

	for scanner.Scan() {
		v := scanner.Text()
		m[v]++
	}
	fmt.Println("word\tcounts")
	for k, v := range m {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	isSHA384 := flag.Bool("sha384", false, "print sha384")
	isSHA512 := flag.Bool("sha512", false, "print sha512")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify args to print sha256")
		os.Exit(1)
	}
	target := []byte(args[0])
	fmt.Printf("%x\n", sha256.Sum256(target))
	if *isSHA384 {
		fmt.Printf("%x\n", sha512.Sum384(target))
	}
	if *isSHA512 {
		fmt.Printf("%x\n", sha512.Sum512(target))
	}
}

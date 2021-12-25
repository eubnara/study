package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

// import "crypto/sha256"

var bitCounts [256]int

func init() {
	for i := 0; i < 256; i++ {
		t := i
		bitCounts[i] = bitCounts[t>>1] + t&1
	}
}

// func countBitsInSHA256(a [32]uint8) int {
// 	counts := int(0)
// 	for _, v := range a {
// 		counts += int(bitCounts[v])
// 	}
// 	return counts
// }

// func realCountBitsInSHA256(a [32]uint8) int {
// 	counts := 0
// 	for _, v := range a {
// 		counts += bits.OnesCount8(v)
// 	}
// 	return counts
// }

func countBitsInUint8(a uint8) int {
	return bitCounts[a]
}

func countDiffInSHA256(a [32]uint8, b [32]uint8) int {
	counts := 0
	for i := 0; i < 32; i++ {
		// counts += countBitsInUint8(a[i] ^ b[i])
		counts += bits.OnesCount8(a[i] ^ b[i])
	}
	return counts
}

func printBits(a [32]uint8) {
	for _, v := range a {
		fmt.Printf("%08b ", v)
	}
	fmt.Println()
}

func main() {
	// fmt.Printf("%T\n", sha256.Sum256([]byte("x")))
	// fmt.Printf("%d\n", countBitsInSHA256(sha256.Sum256([]byte("x"))))
	// fmt.Printf("%d\n", realCountBitsInSHA256(sha256.Sum256([]byte("x"))))
	fmt.Printf("%d\n", countDiffInSHA256(
		sha256.Sum256([]byte("x")),
		sha256.Sum256([]byte("X")),
	))
	// printBits(sha256.Sum256([]byte("x")))
	// printBits(sha256.Sum256([]byte("X")))

	a := byte('x')
	b := []byte("X")
	fmt.Printf("%T %T %d\n", a, b, len(b))
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
}

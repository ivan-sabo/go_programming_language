package main

import (
	"crypto/sha256"
	"fmt"
)

var table [8]byte

func init() {
	for i := 0; i < 8; i++ {
		table[i] = 1 << i
	}
}

func main() {
	a := sha256.Sum256([]byte("a"))
	b := sha256.Sum256([]byte("b"))

	fmt.Printf("%d\n", countDiff(a, b))
}

func countDiff(a, b [32]byte) int {
	diff := 0
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			if (a[i] & table[j]) != (b[i] & table[j]) {
				diff++
			}
		}
	}

	return diff
}

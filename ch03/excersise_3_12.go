// This program checks if two strings are anagrams (contain the same letters in different order).
package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "aabbaa"
	str2 := "ababaa"

	if len(str1) != len(str2) {
		fmt.Println("Strings are not anagagrams")
		return
	}

	r1 := []rune(str1)
	r2 := []rune(str2)

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for i := 0; i < len(r1); i++ {
		m1[r1[i]]++
		m2[r2[i]]++
	}

	if reflect.DeepEqual(m1, m2) {
		fmt.Println("Strings are anagrams")
		return
	}

	fmt.Println("Strings are not anagrams")
}

package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	x = iota + 2
	y
	z
)

func main() {
	invertBits()
	fmt.Println(x)
}

func invertBits() {
	var value uint8 = 113
	fmt.Printf("Binary value: %b\n", value)
	fmt.Printf("Binary value inverted: %b\n", ^value)

}

func convertStringToRune() {
	str := "test"

	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	fmt.Printf("%v\n", n)
	r := rune(n)
	fmt.Printf("%v\n", r)
}

func rawStringLiteral() {
	fmt.Printf(`this is raw string literal
test. everything here
	is interpreted literally, even a new line char \n`)
}

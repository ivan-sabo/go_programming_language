package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("This toll requires exactly one numeric argument")
	}

	x := os.Args[1]
	num, err := strconv.ParseUint(x, 10, 64)
	if err != nil {
		log.Fatalf("An error occured: %e", err)
	}

	fmt.Println(PopCount(num))
}

func PopCount(x uint64) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}

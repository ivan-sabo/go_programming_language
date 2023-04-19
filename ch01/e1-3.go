package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// This is a simple package that prints command line arguments

func main() {
	text := ""
	sep := ""

	start := time.Now()
	for _, v := range os.Args[1:] {
		text += sep + v
		sep = " "
	}
	secs := time.Since(start).Seconds()

	fmt.Println(text)
	fmt.Printf("%.10f\n", secs)

	start = time.Now()
	text = strings.Join(os.Args[1:], " ")
	secs = time.Since(start).Seconds()

	fmt.Println(text)
	fmt.Printf("%.10f\n", secs)
}

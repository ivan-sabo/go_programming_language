package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type info struct {
	count    int
	filename string
}

func main() {
	counts := make(map[string]*info)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if _, ok := counts[line]; !ok {
				counts[line] = &info{
					count:    1,
					filename: filename,
				}
				continue
			}
			counts[line].count++
		}
	}
	for line, info := range counts {
		if info.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", info.count, line, info.filename)
		}
	}
}

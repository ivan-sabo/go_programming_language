// Non-recursive version of comma.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456789"))
}

func comma(s string) string {
	r := []rune(s)
	l := len(r)
	if l <= 3 {
		return s
	}

	var b bytes.Buffer
	for i := 0; i < l; i++ {
		if i != 0 && (l-i)%3 == 0 {
			b.WriteRune(',')
		}
		b.WriteRune(r[i])
	}

	return b.String()
}

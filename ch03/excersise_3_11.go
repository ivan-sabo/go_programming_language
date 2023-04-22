// Non-recursive version of comma, enhanced with floating point and optional sign.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("-12345.6789"))
}

func comma(s string) string {
	r := []rune(s)
	negative := ""
	if r[0] == '-' {
		negative = "-"
		r = r[1:]
	}

	for i, c := range r {
		if c == '.' {
			intPart := integerPart(string(r[:i]))
			fPart := string(r[i+1:])

			return negative + intPart + "." + fPart
		}
	}

	return integerPart(string(r))
}

func integerPart(s string) string {
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

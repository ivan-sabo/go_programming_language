package main

import "fmt"

func main() {
	a := &[4]int{2, 4, 5, 1}
	reverse(a)
	fmt.Println(*a)
}

func reverse(a *[4]int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

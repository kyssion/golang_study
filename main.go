package main

import (
	"fmt"
)

func main() {
	const (
		a, a1 = 1 << iota, 1<<iota - 1
		b, b1
		c, c1
		d     = 123456
		_, d1 = 1 << iota, 1<<iota - 1
		e, e1
	)

	fmt.Println(a, a1)
	fmt.Println(b, b1)
	fmt.Println(c, c1)
	fmt.Println(d, d1)
	fmt.Println(e, e1)
}

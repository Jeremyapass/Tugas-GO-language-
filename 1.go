package main

import "fmt"

func main() {
	var x, a, b, c, d int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		fmt.Scan(&a, &b)
		if a > 0 {
			c += a
		} else {
			d += a
		}
		if b > 0 {
			c += b
		} else {
			d += b
		}
	}
	fmt.Println("Negative: ", d)
	fmt.Println("Positive: ", c)
}

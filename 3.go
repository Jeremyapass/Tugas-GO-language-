package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	konversi(a, &b)
	fmt.Println(b)
}

func pangkat(q, w int) int {
	var i, jum int
	jum = 1
	for i = 1; i <= w; i++ {
		jum *= q
	}
	return jum
}

func konversi(biner int, desimal *int) {
	var a, i int
	*desimal = 0
	i = 0
	for biner > 0 {
		a = biner % 10
		*desimal += a * pangkat(2, i)
		i++
		biner = biner / 10
	}
}

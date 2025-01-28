package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		d = 5
		e = 8
		c = b + a - e*d
	)

	fmt.Println("Nilai C adalah :", c)
}

package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		d = 5
		e = 8
		c = b + a - e*d
		i = 10
		j = 5
	)

	fmt.Println("Nilai C adalah :", c)

	fmt.Println("-------------------Augmented Assigment------------------")
	i += 10
	fmt.Println("Nilai i Sekarang = ", i)

	i += 10
	fmt.Println("Nilai i Sekarang = ", i)

	fmt.Println("-------------------Unary Operator------------------")
	j++
	fmt.Println("Nilai J adalah ", j)
	j++
	fmt.Println("Nilai J++ Sekarang ", j)

	j--
	fmt.Println("Nilai J-- Sekarang ", j)
	j--
	fmt.Println("Nilai J-- Sekarang ", j)
}

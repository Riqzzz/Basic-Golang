package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Mr"
	names[1] = "Towick"
	names[2] = "Aizz"
	fmt.Println(names[0], names[1], names[2])

	var nilai = [3]int{90, 80, 70}
	fmt.Println(nilai)

	var values = [...]int{100, 90, 80}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 120 //mengubah nilai pada indeks 0 menjadi 120
	fmt.Println(values)
}

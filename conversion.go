package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
	)

	fmt.Println(nilai32) //output 32768
	fmt.Println(nilai64) //output = nilai32 karna nilainya dalam jangkauan (reached)
	fmt.Println(nilai16) //output = -32768 karna nilainya tidak dalam jangkauan/max 32767(start dari nilai minimum int16)
}

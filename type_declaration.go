// type declaration digunakan sebagai untuk membuat tipe data baru dari tipe data yg sudah ada/digunakan
// dalam artian terdapat tipe data yang sama dengan sebelumnya dideklarasikan dengan nilai yg ditampung berbeda

package main

import "fmt"

func main() {
	type noTelp string

	var noTowick noTelp = "111111"

	var sample string = "222222"
	var sampleTelp noTelp = noTelp(sample)
	// sampleTlp menggunakan tipedata noTelp pada noTowick sebelumnya dengan output yg ditampung pada tipedata
	// noTelp dari variable sample

	fmt.Println("noTowick akan menampilkan output", noTowick)
	fmt.Println("sampleTelp akan menampilkan output dari sample", sampleTelp)
}

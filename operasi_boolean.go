package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	fmt.Println(lulusNilaiAkhir) //output true

	var lulusAbsensi bool = absensi > 80
	fmt.Println(lulusAbsensi) //output false

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
	// outputnya false karna lulusAbsensi bernilai false pada operator and &&

}

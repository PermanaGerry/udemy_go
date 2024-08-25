package main

import "fmt"

/**
 * Boolean Operator
 * boolean operator adalah operator yang digunakan untuk melakukan operasi logika
 * operator ini menghasilkan nilai boolean yaitu true atau false
 */
func main() {

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 85
	var lulusAbbensi = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbbensi

	fmt.Println("Nilai boolean", lulus)

}

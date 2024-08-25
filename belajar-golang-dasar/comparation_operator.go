package main

import "fmt"

/**
 * Comparation Operator
 * comparation operator adalah operator yang digunakan untuk membandingkan dua buah nilai
 * operator ini menghasilkan nilai boolean yaitu true atau false
 */
func main() {
	// contoh operator perbandingan
	var name1 = "gerry"
	var name2 = "gerry"

	var result bool = name1 == name2
	fmt.Println("Hasil perbandingan", result)

}

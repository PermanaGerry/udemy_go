package main

import "fmt"

/**
 * Array
 * array adalah tipe data yang berisikan beberapa data yang berukuran yang sama
 * contoh deklarasi data array [...]string
 */
func main() {
	// contoh array
	var names [3]string

	names[0] = "Gerry"
	names[1] = "Aria"
	names[2] = "Nutra"

	fmt.Println("Nilai array", names[0])
	fmt.Println("Jumlah array", len(names))

	// contoh array shorcut
	var names2 = [3]string{"Gerry", "Aria", "Nutra"}

	fmt.Println("Nilai array", names2[2])
	fmt.Println("Jumlah array", len(names2))

}

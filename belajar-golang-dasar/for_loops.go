package main

import "fmt"

/**
 * For Loops
 * For loops adalah sebuah syntax yang digunakan untuk melakukan perulangan
 */
func main() {

	people := []string{"Gerry", "Aria", "Nutra"}

	// contoh for
	fmt.Println("-- contoh for --")

	var i = 0
	for i < len(people) {
		fmt.Println(people[i])
		i++
	}

	/**
	 * for statement
	 * Dalam for kita bisa menambahkan 3 statement yaitu :
	 * 1. statement init (sebelum for di eksekusi)
	 * 2. statement condition (kondisi di eksekusi)
	 * 3. statement post (setelah for di eksekusi)
	 * syntax : for init; condition; post {
	 * }
	 */

	// contoh for dengan statement
	fmt.Println("-- contoh for dengan statement --")
	for i := 0; i < len(people); i++ {
		fmt.Println("", people[i])
	}

	/**
	 * range statement
	 * range statement adalah statement yang digunakan untuk melakukan perulangan
	 * range statement memiliki 2 buah operand yaitu array dan slice
	 */

	// contoh range statement
	fmt.Println("-- contoh range statement --")
	for index, value := range people {
		fmt.Println("index", index, "value", value)
	}

}

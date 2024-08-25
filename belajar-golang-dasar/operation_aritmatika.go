package main

import "fmt"

/**
 * Operation Aritmatika
 * operator aritmatika adalah operator yang digunakan untuk melakukan operasi matematika
 * seperti penjumlahan (+), pengurangan (-), perkalian (*), pembagian (/), dan modulus (%)
 */
func main() {
	// contoh penjumlahan
	var a = 10
	var b = 10
	var c = 20

	count := (a + b) - c
	fmt.Println("Hasil penjumlahan", count)

	/**
	 * augmented assignment
	 * augmented assignment adalah cara singkat untuk melakukan operasi aritmatika
	 */

	// contoh modulo
	var i = 20
	i %= 18

	fmt.Println("Hasil modulo", i)

	/**
	 * unary operator
	 * unary operator adalah operator yang dilakukan pada satu operand
	 * contoh unary operator adalah operator increment (++) dan decrement (--)
	 */

	// contoh increment
	var j = 10
	j++

	fmt.Println("Hasil unary", j)

}

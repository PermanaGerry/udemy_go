package main

import "fmt"

/**
 * if expression
 * if expression adalah statement yang digunakan untuk mengeksekusi statement jika sebuah kondisi bernilai true
 * contoh if expression
 */
func main() {
	name := "Nutra"

	// contoh if else
	if name == "Aria" {
		fmt.Println("Hello Aria")
	} else {
		fmt.Println("Benar nama anda adalah", name)
	}

	// contoh if else if
	if name == "Aria" {
		fmt.Println("Hello Aria")
	} else if name == "Nutra" {
		fmt.Println("Hello Nutra")
	} else {
		fmt.Println("Benar nama anda adalah", name)
	}

	// contoh if short
	if length := len(name); length > 5 {
		fmt.Println("contoh if short", length)
	} else {
		fmt.Println("contoh else short", length)
	}
}

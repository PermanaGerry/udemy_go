package main

import "fmt"

/**
 * switch expression
 * switch expression adalah statement yang digunakan untuk mengeksekusi statement jika sebuah kondisi bernilai true
 */
func main() {

	// contoh switch expression
	var name = "Gerry"
	switch name {
	case "Aria":
		fmt.Println("Hello Aria")
	case "Nutra":
		fmt.Println("Hello Nutra")
	default:
		fmt.Println("Benar nama anda adalah", name)
	}

	// contoh switch expression short
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("contoh switch expression short", lenght)
	case false:
		fmt.Println("contoh else switch expression short", lenght)
	}

	// contoh switch expression fallthrough
	switch {
	case name == "Aria":
		fmt.Println("Hello Aria")
		fallthrough
	case name == "Nutra":
		fmt.Println("Hello Nutra")
	default:
		fmt.Println("Benar nama anda adalah", name)
	}
}

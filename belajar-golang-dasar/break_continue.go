package main

import "fmt"

/**
 * break and continue
 * break dan continue adalah statement yang digunakan untuk menghentikan perulangan atau melanjutkan perulangan
 */
func main() {
	// contoh break
	fmt.Println("-- contoh break --")
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan", i)

		if i == 1 {
			break
		}
	}

	// contoh continue
	fmt.Println("-- contoh continue --")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan", i)
	}
}

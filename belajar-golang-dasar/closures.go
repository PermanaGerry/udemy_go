package main

import "fmt"

/**
 * Closure
 * Closure adalah kemampuan sebuah function berinteraksi
 * dengan data-data yang ada di luar function
 */
func main() {
	counter := 1
	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}

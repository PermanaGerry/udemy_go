package main

import "fmt"

func data() interface{} {
	return 100
}

/**
 * Type Assertions
 * type assertions adalah fitur dalam golang yang memungkinkan kita untuk melakukan casting tipe data
 */
func main() {

	// contoh type assertions
	result := data()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

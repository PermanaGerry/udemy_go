package main

import "fmt"

/**
 * Function
 * function adalah blok kode yang dapat dipanggil kembali
 */
func main() {

	fmt.Println("-- Function tidak mengembalikan nilai dan tidak menerima parameter --")
	hello()

	fmt.Println("-- Function tidak mengembalikan nilai dan menerima parameter --")
	hello2("Gerry")

	fmt.Println("-- Function mengembalikan nilai dan menerima parameter --")
	result := hello3("Gerry")
	fmt.Println(result)

	fmt.Println("-- Function mengembalikan nilai lebih dari satu parameter --")
	resultFirstName, resultMidleName, resultLastName := hello4("Gerry", "Nutra", "Aria")
	fmt.Println(resultFirstName)
	fmt.Println(resultMidleName)
	fmt.Println(resultLastName)

	fmt.Println("-- Untuk menghiraukan pengembalian salah satu parameter --")
	resultFirstName2, _, _ := hello4("Gerry", "Nutra", "Aria")
	fmt.Println(resultFirstName2)

	fmt.Println("-- Function variadic --")
	resultVariadic := hello5("Gerry", "Nutra", "Aria")
	fmt.Println(resultVariadic)
}

/**
 * function tidak mengembalikan nilai dan tidak menerima parameter
 */
func hello() {
	fmt.Println("Hello World")
}

/**
 * function tidak mengembalikan nilai dan menerima parameter
 */
func hello2(name string) {
	fmt.Println("Hello", name)
}

/**
 * function mengembalikan nilai dan menerima parameter
 */
func hello3(name string) string {
	return "Hello " + name
}

/**
 * function mengembalikan nilai lebih dari satu parameter
 */
func hello4(firstName string, midleName string, lastName string) (string, string, string) {
	resultFirstName := "first name: " + firstName
	resultMidleName := "middle name: " + midleName
	resultLastName := "last name: " + lastName

	return resultFirstName, resultMidleName, resultLastName
}

/**
 * function variadic atau pembuatan fungsi dengan parameter bisa
 * menampung nilai sejenis yang tidak terbatas jumlahnya.
 */
func hello5(names ...string) string {
	var result string

	for _, name := range names {
		result += name + " "
	}

	return result
}

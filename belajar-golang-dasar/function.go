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
	resultFirstName, resultMiddleName, resultLastName := hello4("Gerry", "Nutra", "Aria")
	fmt.Println(resultFirstName)
	fmt.Println(resultMiddleName)
	fmt.Println(resultLastName)

	fmt.Println("-- Untuk menghiraukan pengembalian salah satu parameter --")
	resultFirstName2, _, _ := hello4("Gerry", "Nutra", "Aria")
	fmt.Println(resultFirstName2)

	fmt.Println("-- Function variadic --")
	resultVariadic := hello5("Gerry", "Nutra", "Aria")
	fmt.Println(resultVariadic)

	fmt.Println("-- Function pengembalian default --")
	resultFirstName1, resultMiddleName1, _ := hello6("Gerry", "Nutra", "Aria")
	fmt.Println(resultFirstName1)
	fmt.Println(resultMiddleName1)
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
func hello4(firstName string, middleName string, lastName string) (string, string, string) {
	resultFirstName := "first name: " + firstName
	resultMiddleName := "middle name: " + middleName
	resultLastName := "last name: " + lastName

	return resultFirstName, resultMiddleName, resultLastName
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

/**
 * Function pengembalian default
 */
func hello6(firstName string, middleName string, lastName string) (resultFirstName, resultMiddleName, resultLastName string) {
	resultFirstName = "first name: " + firstName
	resultLastName = "last name: " + lastName

	return resultFirstName, resultMiddleName, resultLastName
}

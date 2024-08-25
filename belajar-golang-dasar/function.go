package main

import (
	"fmt"
)

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

	fmt.Println("-- Function variadic dengan data slice --")
	namesSlice := []string{"Gerry", "Nutra", "Aria"}
	resultNamesSlice := hello5(namesSlice...)
	fmt.Println(resultNamesSlice)

	fmt.Println("-- Function pengembalian default --")
	resultFirstName1, resultMiddleName1, _ := hello6("Gerry", "Nutra", "Aria")
	fmt.Println(resultFirstName1)
	fmt.Println(resultMiddleName1)

	fmt.Println("-- Function digunakan sebagai value --")
	funcAsValue := hello2
	funcAsValue("Nutra")

	fmt.Println("-- Function digunakan sebagai parameter --")
	hello7("Gerry", filterName)

	filter := filterName
	hello7("Nutra", filter)

	fmt.Println("-- Function type declaration --")
	filterNameAge := filterNameAge
	hello8("Gerry", 23, filterNameAge)
	hello8("Nutra", 23, filterNameAge)

	fmt.Println("-- Function anonymous --")
	funcAnonymous := func(name string) bool {
		return name == "kuda"
	}
	fmt.Println(funcAnonymous("kuda"))
	fmt.Println(funcAnonymous("kucing"))

	fmt.Println("-- Function recursive --")
	fmt.Println(factorial(10))
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

/**
 * Function digunakan sebagai parameter
 */
func hello7(name string, filter func(string) string) {
	fmt.Println("Hello", name)
}

func filterName(name string) string {
	if name == "Gerry" {
		return "Aria"
	}

	return name
}

/**
 * Function type declaration
 */

type filterAge func(string, int) bool

func hello8(name string, age int, filter filterAge) {
	if filter(name, age) {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Your age is", age)
	}
}

func filterNameAge(name string, age int) bool {
	if name == "Gerry" && age == 23 {
		return true
	} else {
		return false
	}
}

/**
 * Function recursive
 */

func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

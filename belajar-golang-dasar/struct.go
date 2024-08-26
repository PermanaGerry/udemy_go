package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int8
}

/**
 * Struct
 * Struct adalah sebuah template data yang digunakan untuk menampung data
 */
func main() {

	// contoh struct
	fmt.Println("-- contoh struct --")
	var customer Customer

	customer.Name = "Gerry Aria Nutra"
	customer.Address = "Jl. Cempaka No. 1"
	customer.Age = 23

	fmt.Println(customer)

	// contoh struct literal
	fmt.Println("-- contoh struct literal --")
	customer2 := Customer{
		Name:    "Gerry Aria Nutra",
		Address: "Jl. Cempaka No. 1",
		Age:     23,
	}

	fmt.Println(customer2)

	customer3 := Customer{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}
	fmt.Println(customer3)

	// contoh struct method
	fmt.Println("-- contoh struct method --")
	customer.sayHello("Utomo")
}

/**
 * struct method
 * struct method adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
 */
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

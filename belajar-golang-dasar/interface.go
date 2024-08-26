package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

/**
 * Interface kosong
 * Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat
 * secara otomatis semua tipe data akan menjadi implementasinya
 */
func EmptyInterface(value interface{}) interface{} {
	return value
}

/**
 * Interface
 * Interface adalah tipe data yang berisi definisi method
 * Interface biasanya digunakan untuk kontrak
 */
func main() {

	// contoh interface
	fmt.Println("-- Contoh Interface --")
	person := Person{"Gerry Aria Nutra"}
	SayHello(person)

	// contoh interface kosong
	fmt.Println("-- Contoh Interface Kosong --")
	fmt.Println(EmptyInterface("Hello"))
}

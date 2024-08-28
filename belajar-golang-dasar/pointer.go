package main

import "fmt"

type Person struct {
	Name    string
	Address string
	Age     int
}

/**
 * Pointer
 * pointer adalah sebuah tipe data yang digunakan untuk menyimpan alamat memory dalam sebuah variable
 * untuk mendeklarasikan pointer pada varibel harus di awali dengan &(dan) / *(bintang)
 * contoh
 * - dengan menggunakan tanda *(bintang)
 * type Persen *[]string
 * - dengan menggunakan tanda &(dan)
 * var person2 = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}
 * var replaceName = &person2
 */
func main() {

	// contoh pointer pass by value
	fmt.Println("-- contoh pass by value --")
	passByValue()

	// contoh pointer pas by refrence
	fmt.Println("-- contoh pointer pas by refrence --")
	pointerPassByRefrence()
}

/**
 * pass by value
 * Artinya data yang ada di varibel baru merupakan data duplikat dari varibel sebelumnya,
 * jika data dirubah data sebelumnya tidak berubah
 *
 * note : dalam golang untuk default value sebuah variabel atau fungsi adalah pass by value
 */
func passByValue() {
	var person = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}

	var copyPerson = person
	copyPerson.Name = "Aria"

	fmt.Println(person, copyPerson)
}

/**
 * pass by refrence
 * Artinya data yang ada di dalam variabel merupakan reference data dari sebelumnya
 *
 * note : untuk merubah data sebelumnya harus menggunakan pointer
 */
func pointerPassByRefrence() {
	var person = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}

	var replaceName = &person
	replaceName.Name = "Aria"

	fmt.Println(person, replaceName)
}

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
 * Ada 2 hal penting yang perlu diketauhui mengenai pointer :
 * - untuk mengambil nilai pointer bisa menggunakan tanda ampersand (&) metode ini disebut refrence
 * contoh
 * var person = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}
 * var person1 = &person
 * - untuk mengambil nilai asli sebuah pointer bisa menggunakan tanda asterisk (*) metode ini disebut dereferencing
 * contoh
 * var person = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}
 * var person1 *Person = &person
 */
func main() {

	// contoh pass by value
	fmt.Println("-- contoh pass by value --")
	passByValue()

	// contoh pass by refrence
	fmt.Println("-- contoh pass by refrence --")
	passByRefrence()

	// contoh swap value by refrence dengan menggunakan operator arsterisk(*)
	fmt.Println("-- contoh swap value by refrence --")
	swapValueByRefrence()
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
func passByRefrence() {
	var person Person = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}

	var replaceName *Person = &person
	replaceName.Name = "Aria"

	fmt.Println(person, replaceName)
}

/**
 * swap value by refrence
 * Artinya mengganti data refrence dan data asli / data sebelumnya ke data yang baru
 * untuk mengganti value bisa menggunakan operator asterisk(*)
 */
func swapValueByRefrence() {
	var person Person = Person{"Gerry Aria Nutra", "Jl. Cempaka No. 1", 23}

	var replacePersonName *Person = &person

	replacePersonName.Name = "Aria"

	fmt.Println("Data yang diganti pada key nama", *replacePersonName)

	*replacePersonName = Person{"Nagato Uzumaki", "Jl. Bandung No. 1", 45}

	fmt.Println("Data yang sudah di ganti pada varibel person dan replacePersonName dengan operator arsterisk")
	fmt.Println("person : ", person)
	fmt.Println("replacePersonName : ", *replacePersonName)
}

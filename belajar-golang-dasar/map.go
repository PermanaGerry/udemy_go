package main

import "fmt"

/**
 * map
 * map adalah tipe data khusus yang berisikan key dan value
 * contoh deklarasi data map map[string]string
 */
func main() {
	// contoh map
	people := map[string]string{
		"name":    "Gerry Aria Nutra",
		"age":     "23",
		"address": "Jl. Cempaka No. 1",
	}

	// menampilkan semua data
	fmt.Println(people)

	// menampilkan data berdasarkan key
	fmt.Println(people["name"])

	/**
	 * function map
	 * yang terdapat pada function map adalah len, delete dan make
	 */

	// contoh make
	newPeople := make(map[string]string)
	newPeople["people1"] = "Gerry"
	newPeople["people2"] = "Aria"
	newPeople["people3"] = "Nutra"

	// contoh len
	fmt.Println("nilai len", len(newPeople))

	// contoh delete
	delete(newPeople, "people1")
	fmt.Println("nilai delete", newPeople)
}

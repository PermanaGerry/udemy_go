package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

/**
 * Nil
 * nil adalah sebuah data kosong hanya bisa digunakan pada beberapa tipe data,
 * yaitu interface, map, slice, pointer, function, dan channel
 */
func main() {
	// contoh nil
	var person = NewMap("Nutra")
	if person == nil {
		fmt.Println("Data person kosong")
	} else {
		fmt.Println("Data person ada:", person)
	}
}

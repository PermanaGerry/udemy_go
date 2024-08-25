package main

import "fmt"

/**
 * slice
 * slice adalah tipe data array yang berisikan beberapa data yang berukuran yang berbeda
 * contoh deklarasi data slice []string
 */
func main() {
	// contoh slice
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println("nilai slice", slice1)

	/**
	 * function slice
	 * yang terdapat pada function slice adalah len, cap, append, make, copy
	 */

	// contoh len
	fmt.Println("nilai jumlah data", len(slice1))
	// contoh cap
	fmt.Println("nilai kapasitas", cap(slice1))

	// contoh append
	var slice2 = append(slice1, "Romeo")
	fmt.Println("nilai slice", slice2)
	fmt.Println("Jumlah data", len(slice2))
	fmt.Println("nilai kapasitas", cap(slice2))

	// contoh make
	var days = make([]string, 7, 7)
	days[0] = "minggu"
	days[1] = "senin"
	days[2] = "selasa"
	days[3] = "rabu"
	days[4] = "kamis"
	days[5] = "jumat"
	days[6] = "sabtu"

	fmt.Println("Nilai slice", days)
	fmt.Println("Jumlah data", len(days))
	fmt.Println("Nilai kapasitas", cap(days))

	// contoh new slice in make
	newDays := append(days, "libur")
	fmt.Println("Nilai slice", newDays)
	fmt.Println("Jumlah data", len(newDays))
	fmt.Println("Nilai kapasitas", cap(newDays))

}

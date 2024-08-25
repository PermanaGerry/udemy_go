package main

import "fmt"

/**
 * defer
 * defer adalah statement yang dapat digunakan bila applikasi mengalami error
 *
 * panic
 * panic adalah statement digunakan untuk menghentikan sebuah aplikasi dengan memberikan sebuah pesan error
 *
 * recover
 * recover adalah statement yang digunakan untuk mengembalikan sebuah error yang terjadi dan mengembalikan sebuah value
 * kemudian aplikasi berjalan dengan normal
 */
func main() {
	runApp(true)

	fmt.Println("The application runs with normal flow even though it panics")
}

func endApp() {
	fmt.Println("End Application")

	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application Error !!!")
	}

}

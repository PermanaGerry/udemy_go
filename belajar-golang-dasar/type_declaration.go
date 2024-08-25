package main

import "fmt"

func main() {
	type noKtp string

	var noKtpEko noKtp = "1234567890"

	var noKtpGerry string = "0987654321"
	var noKtpGerryAlias noKtp = noKtp(noKtpGerry)

	fmt.Println(noKtpEko)
	fmt.Println(noKtpGerryAlias)
}

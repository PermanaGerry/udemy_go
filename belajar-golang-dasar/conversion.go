package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)
	fmt.Println(value8)

	var name = "Gerry permana putra"
	var e uint8 = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}

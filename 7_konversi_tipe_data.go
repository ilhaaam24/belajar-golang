package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	// hasilnya -32768 karena melebihi batas dan kembali ke paling bawah yaitu -32768 atau bisa disebut number overflow
	fmt.Println(nilai16)




	var name = "Muhammad Ilham"
	var m  uint8= name[0]
	var stringM = string(m)
	fmt.Println(m)
	fmt.Println(stringM)
}
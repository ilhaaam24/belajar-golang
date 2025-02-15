package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {


	// pada address 1 tidak berubah karena address 2 mengcopy value dari address, atau bisa disebut pass by value
	// address := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := address

	// address2.City = "Surabaya"

	// fmt.Println(address)
	// fmt.Println(address2)


	// pada address 1 ikut berubah karena address 2 mengcopy reference dari address, atau bisa disebut pass by reference
	address := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address

	address2.City = "Surabaya"

	fmt.Println(address)
	fmt.Println(address2)

}

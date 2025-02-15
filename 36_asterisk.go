package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address

	address2.City = "Surabaya"

	fmt.Println(address)
	fmt.Println(address2)


	// address 1 tidak ikut berubah karena address 2 merujuk ke address baru
	// address2 = &Address{" Malang", "DKI Jakarta", "Indonesia"}

//  address 1 ikut berubah karena address 2 merujuk ke address yang sama dan membuat address baru
	*address2 = Address{" Malang", "DKI Jakarta", "Indonesia"}

	fmt.Println(address)
	fmt.Println(address2)	
}

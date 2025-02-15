package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// ketika menggunakan new maka var address otomatis tipe datanya pointer
	address := new(Address)
	changeCountryToIndonesia(address)

	fmt.Println(address)

}
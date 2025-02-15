package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {


	//  new  digunakan ketika data masih kosong dan ingin membuat pointer
	var alamat1  *Address = new(Address)
	var alamat2 *Address = alamat1

	// alamat 1 dan alamat 2 akan merujuk ke alamat yang sama yaitu &Address

	alamat2.City = "Surabaya"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}

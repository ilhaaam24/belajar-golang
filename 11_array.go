package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Ilham"
	names[2] = "Budi"
	// names[3] = "Joko" // error karena array hanya memiliki 3 index
	// names[3] = "Joko"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi array langsung 
	var values = [3]int{
		60,
		10,
		// data ketiga default 0
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// function array =
	// len mengembalikan panjang array
	// array[index] mengembalikan data array sesuai index
	//array[index] = value mengubah data array sesuai index


	// ... adalah digunakan untuk menginitialisasi array
	var values2 = [...]int{
		60,
		10,
		70,
	}

	fmt.Println(len(values2))

	fmt.Println(values2[0])
	values2[2] = 100
	fmt.Println(values2[2])
}
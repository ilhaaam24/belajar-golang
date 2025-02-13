package main

import "fmt"

func main(){

	// map adalah tipe data yang digunakan untuk menyimpan data dengan key dan value
	names := map[string]string{
		"joko": "joko",
		"budi": "budi",
		"eko": "eko",
	}

	names["kurniawan"] = "kurniawan"

	for _,name := range names {
		fmt.Println(name)
	}

}
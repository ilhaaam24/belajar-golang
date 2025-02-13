package main

import "fmt"

func main() {

	names := [...]string{"joko", "budi", "eko", "kurniawan"}

	names2 := names[1:3]
	names2[0] = "ilham"

	fmt.Println(names)
	fmt.Println(names2)
	fmt.Println(len(names2))
	fmt.Println(cap(names2))
	names2 = append(names2, "jokowi")

	fmt.Println(names)
	fmt.Println(names2)
	
	// jika kapasitas pada slice melebihi kapasitas array maka akan membuat array baru dan tidak akan merupkan array lama
	names2 = append(names2, "widodo")
	fmt.Println(names)
	fmt.Println(names2)
	names2[0]=" prabowo"	
	
	fmt.Println(names)
	fmt.Println(names2)
	


}
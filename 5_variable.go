package main

import "fmt"

func main(){

	var name string

	name = "Muhammad Ilham"
	fmt.Println(name)
	
	
	name = "Ilham"
	fmt.Println(name)

// tidak perlu deklarasi tipe data
	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// tidak perlu deklarasi var denga titik 2, hanya untuk deklarasi awal
	friendAge := 30
	fmt.Println(friendAge)


	Country := "Indonesia"
	fmt.Println(Country)



	// multiple variable
	var(
		firstName = "Muhammad" 
		lastName = "Ilham"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
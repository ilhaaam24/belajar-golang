package main

import "fmt"

func helloName() (firstName, middleName, lastName string) { // named return value
	firstName = "Muhammad"
	middleName = "Ilham"
	lastName = "Budi"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, _ := helloName()
	fmt.Println(firstName, middleName)

}
package main

import "fmt"

type User struct {
	Name string
}


// pointer penting digunakan ketika memanipulasi data struct menggunakan method
func (user *User) Married() {
	user.Name = "Mr " + user.Name
}

func main() {
	ilham := User{"Ilham"}
	ilham.Married()

	fmt.Println(ilham.Name)
}
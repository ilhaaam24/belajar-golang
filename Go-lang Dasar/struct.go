package main

import "fmt"

type User struct {
	Name, Address string
	Age           int
}

type HasName interface{
	GetName() string
}

func SayHello(hasName HasName ){
	fmt.Println("Hello", hasName.GetName())
}

func (user User) GetName() string{
	return user.Name
} 


func main() {
	user := User{"Joko", "Jakarta", 20}
	fmt.Println(user)

	SayHello(user)
	
}
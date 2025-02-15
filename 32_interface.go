package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHello(value HashName) {
	fmt.Println("Hello", value.GetName())

}


type Person struct{
	Name string
}


func (person Person) GetName() string{
	return person.Name
}

func main() {
	person := Person{"ilham"}


	SayHello(person)
}
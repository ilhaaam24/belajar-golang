package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}


func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name , "My Name is", customer.Name)
}

func main() {

	var eko Customer
	eko.Name = "Eko Krniawan"
	eko.Address = "Jakarta"
	eko.Age = 30
	fmt.Println(eko)


	 budi := Customer{"budi", "Surabaya", 20}

	 fmt.Println(budi)
	 fmt.Println(budi.Name)
	 fmt.Println(budi.Address)
	 fmt.Println(budi.Age)
	 
	
	 budi.sayHello("Ilham")
 }
package main

import "fmt"

// function sebagai parameter type declaration
type Filter func(string) string

func sayHello2(name string, filter Filter) {
	filteredName := filter(name)

	fmt.Println("Hello", filteredName)
}



// function sebagai parameter standard
func sayHello(name string, filter func(string) string) {
	filteredName := filter(name)

	fmt.Println("Hello", filteredName)
}


func filterName(name string) string{

	if name == "Anjing"{
		return "..."

	}else{
		return name
	}
}

func main() {

	sayHello("Anjing", filterName)



	sayHello2("Ilham", filterName)
}
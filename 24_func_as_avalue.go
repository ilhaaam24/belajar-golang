package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	/**
	 * func sebagai value
	 */

	// func sebagai value , dan contoh menjadi function
	contoh := sayHello

	fmt.Println(contoh("Ilham"))
	
}
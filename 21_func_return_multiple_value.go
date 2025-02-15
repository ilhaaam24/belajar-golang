package main

import "fmt"

func sayHello(name string) ( string, string) {
	return "Hello", name
}

func main() {

	// text hello sebagai pertama dan name sebagai kedua
	hello, name := sayHello("Ilham")
	fmt.Println(hello, name)


	//  _ digunakan jika tidak ingin mengambil value pertama
	_, name = sayHello("Joko")
	fmt.Println(name)
}
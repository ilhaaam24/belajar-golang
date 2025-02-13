package main

import "fmt"

func main() {
	// operasi boolean && || !

	var data1 bool = true
	var data2 bool = false

	fmt.Println(data1 && data2) // false
	fmt.Println(data1 || data2) // true
	fmt.Println(!data1)         // false

}
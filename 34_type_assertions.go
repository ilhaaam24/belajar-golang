package main

import "fmt"

func result() interface{} {
	return true
}

func main() {

	value := result()
	switch value := value.(type) {
	case string:
		fmt.Println("string", value)	
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)

	}
}
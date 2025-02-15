package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)


func Hallo(name string) string {
	return "Hello " + name
}

func main(){
	result := helper.SayHello("Eko")
	fmt.Println(result)

}
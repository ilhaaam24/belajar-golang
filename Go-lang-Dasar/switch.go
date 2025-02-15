package main

import "fmt"

func main(){
	number := 5
 switch number {
 case 5:
		fmt.Println("number is less than 5")
 case 8:
		fmt.Println("number is less than 8")
 default:
	fmt.Println("number is greater than 8")
 }
}
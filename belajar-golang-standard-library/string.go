package main

import (
	"fmt"
	"strings"
)



func main(){
	fmt.Println(strings.Contains("Muhammad Ilham", "Ilham"))
	fmt.Println(strings.Split("Muhammad Ilham", " "))
	fmt.Println(strings.ToLower("Muhammad Ilham"))
	fmt.Println(strings.ToUpper("Muhammad Ilham"))
	fmt.Println(strings.Trim("    Muhammad Ilham    ", " "))
	fmt.Println(strings.ReplaceAll("Muhammad Ilham", "Ilham", "Iqbal"))
	fmt.Println(strings.Repeat(" Muhammad Ilham",5))
	fmt.Println(strings.ToTitle(" Muhammad Ilham")) //change the string to title uppercase

}
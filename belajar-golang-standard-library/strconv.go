package main

import (
	"fmt"
	"strconv"
)


func main(){
	result, err := strconv.Atoi("100")
	if err == nil{
		fmt.Println(result)
	}else{
		fmt.Println("Error", err.Error())
	}


	strconvBool, err := strconv.ParseBool("true")
	if err == nil{
		fmt.Println(strconvBool)
	}else{
		fmt.Println("Error", err.Error())
	}


	fmt.Println(strconv.Itoa(100))
	fmt.Println(strconv.FormatInt(100, 2)) // binary
}
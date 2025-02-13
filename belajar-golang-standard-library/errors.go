package main

import (
	"errors"
	"fmt"
)



var(
	validationError = errors.New("validation error")
	notFoundError = errors.New("not found error")
)



func GetById(name string) error{
	if name == ""{
		return validationError
	}else if name != "eko"{
		return notFoundError
	}



	return nil
}

func main() {
	err := GetById("eko")
	if err != nil{
		if errors.Is(err, validationError){
			fmt.Println("validation error")
		}else if errors.Is(err, notFoundError){
			fmt.Println("not found error")
		}else{
			fmt.Println("unknown error")
		}
		
	}else{
		fmt.Println("success")
	}



}
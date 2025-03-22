package main

import (
	"fmt"
	"reflect"
)


type Person struct{
	Name, Address, Email string
}

type Sample struct{
	Name string 
}


func readField (value any){
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type : ", valueType.Name())


	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Printf("Type: %v, Name: %v\n", valueField.Type, valueField.Name) 
	}
}



func main(){
	var person = Person{Name: "Rizky", Address: "Bandung", Email: "2h4eM@example.com"}
	readField(person)
}
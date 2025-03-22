package main

import (
	"fmt"
	"reflect"
)


type Person struct{
	Name string `required:"true" max:"10"`
	 Address string `required:"true"max:"10"`
	  Email string `required:"true"max:"10"`
}

type Sample struct{
	Name string `required:"true"`
}


func readField (value any){
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type : ", valueType.Name())


	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Printf("Type: %v, Name: %v\n", valueField.Type, valueField.Name ) 
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}
}


func isValid(value any)(result bool){
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true"{
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false{
				return false
			}
		}
		
	}
	return result
}


func main(){
	var person = Person{Name: "Rizky", Address: "", Email: "2h4eM@example.com"}
	readField(person)
	fmt.Println("result : ",isValid(person)) // false karena address masih kosong
}
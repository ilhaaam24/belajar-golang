package main

import "fmt"

// nil hanya dapat digunakan pada tipe data pointer, channel, map, slice, dan function

// error
// func Contoh( name string) string{
// 	if name == ""{
// 		return nil
// 	}
// }

func NewMap(name string) map[string]string {
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name" : name,
		}
	}
}


func main(){

data := NewMap("Ilham")

if data == nil{
	fmt.Println("Data Kosong")
}else{
	fmt.Println(data["name"])}


}
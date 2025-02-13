package main

import "fmt"


func logging (){
 fmt.Println("selesai memanggil function")
}

func endApp(){
 fmt.Println("Aplikasi Selesai")	
 message := recover()
 fmt.Println("terjadi error", message)
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Error pada bagian 	aplikasi")
	}
}
func main(){
 runApp(true)
}
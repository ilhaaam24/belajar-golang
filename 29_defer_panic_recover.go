package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
	// recover digunakan untuk menambahkan penanganan error
	message:= recover()
	fmt.Println("iyaaaa")
	fmt.Println("terjadi error", message)
}

func runApp(error bool){
	// defer akan tetap dijalankan meskipun terjadi error
	defer logging()
	if error{
		panic("Error pada bagian aplikasi")
	}


	fmt.Println("halooooo")
}

func main() {
runApp(true)
}
package main

import "fmt"

func tambah(a int, b int) int {
	return a + b
}

func main() {
	sum := tambah(2, 3)
	fmt.Println(sum)
}
package main

import "fmt"

func main() {

	// operator matamatika + - * : / %

	var a int = 10
	var b int = 5

	var c int = a + b
	var d int = a - b
	var e int = a * b
	var f int = a / b
	var g int = a % b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// augmented assignment += -= *= /= %=

	var x int = 10
	x += 10
	fmt.Println(x)

	var y int = 10
	y -= 10
	fmt.Println(y)

	// unary operator ++ --


	// hanya bertambah 1
	var z int = 10
	z++
	fmt.Println(z)

	var w int = 10
	w--
	fmt.Println(w)


	// operasi perbandingan > < >= <= == != yang akan menghasilkan boolean

	var h int = 10
	var i int = 20

	var hasil bool = h < i
	fmt.Println(hasil)
}
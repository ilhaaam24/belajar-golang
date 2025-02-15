package main

import "fmt"

func factorial(n int) int {
	value := 1
	for i := n; i > 0; i-- {
		value *= i
	}
	return value
}


func recursiveFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

func main() {
	fmt.Println(10*9*8*7*6*5*4*3*2*1)
	fmt.Println(factorial(10))

	fmt.Println(recursiveFactorial(10))
}
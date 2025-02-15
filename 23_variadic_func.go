package main

import "fmt"

func sumAll(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}

func main() {

	fmt.Println(sumAll(1, 2, 3, 4, 5))
}
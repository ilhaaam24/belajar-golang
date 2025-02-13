package main

import "fmt"

func main() {

	// perulangan ini akan hentikan ketika i = 5 dan tidak akan melanjutkan perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	// perulangan ini akan melanjutkan perulangan selanjutnya ketika i = 5 tanpa print 5
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
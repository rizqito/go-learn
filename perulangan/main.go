package main

import "fmt"

func main() {
	// perulangan for biasa
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// perulangan for while
	var j = 0
	for {
		fmt.Println("Angka", j)
		j++
		if j == 5 {
			break
		}
	}

	// perulangan for range
	fruits := []string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}

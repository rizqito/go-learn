package main

import "fmt"

func arrayEx() {
	var names [4]string
	names[0] = "asdasd"
	names[1] = "asdasd"
	names[2] = "asdasd"
	names[3] = "asdasd"

	var fruits = [4]string{"apple", "grape", ""}

	fmt.Println("jumlah data names ", len(names))
	fmt.Println("jumlah data fruits", len(fruits))

	for _, value := range names {
		fmt.Println(value)
	}

	for _, value := range fruits {
		fmt.Println(value)
	}
}

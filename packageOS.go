package main

import (
	"fmt"
	"os"
)

func osPackage() {
	args := os.Args
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println("Error", err.Error())
	}
}

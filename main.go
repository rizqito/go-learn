package main

import (
	"fmt"
	"go-learn/helper"
)

func main() {
	// cara pemanggilan function dalam package lain

	fmt.Println("-----Hello World-------")
	helper.SayHello("Rizqi")
}

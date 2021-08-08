package main

import (
	"fmt"
	"go-learn/helper"
)

func main() {
	fmt.Println("-----Hello World-------")
	helper.SayHello("Rizqi")

	fmt.Println("------variabel---------")
	helper.Variabel()

	fmt.Println("------Try--------------")
	helper.Try()

	fmt.Println("------Array Slice------")
	helper.ArraySlices()

	fmt.Println("-------Map-------------")
	helper.Map()

	fmt.Println("-------Pointer---------")
	helper.Pointer()

	fmt.Println("-------PrintF----------")
	helper.CetakF()

	fmt.Println("-------Random----------")
	helper.Random()

	fmt.Println("-------SQRT------------")
	helper.Sqrt()

	fmt.Println("-------Structs---------")
	helper.BelajarStruct()

	fmt.Println("-------Function--------")
	helper.Fungsi()

	fmt.Println("--------Method---------")
	helper.BelajarMethod()

	fmt.Println("-------INterface-------")
	helper.Interface()
}

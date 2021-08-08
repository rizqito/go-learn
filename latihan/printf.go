package main

import "fmt"

func main(){
	i, f, g := 42, 3.142, 0.867 + 0.5i
	fmt.Printf("i mempunyai tipe data %T", i)
	fmt.Println()
	fmt.Printf("f mempunyai tipe data %T", f)
	fmt.Println()
	fmt.Printf("g mempunyai tipe data %T", g)
	fmt.Println()
}
package main

import "fmt"

func main() {

	//? variabel beserta tipe data
	a := "Hello a\n" //* penggunaan variabel shorthand
	var b string     //* penggunaan variabel biasa
	b = "Hello b\n"
	var c = "Hello c\n"        //* penggunaan variabel biasa
	var d string = "Hello d\n" //* penggunaan variabel biasa
	fmt.Println(a, b, c, d)
}

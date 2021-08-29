package main

import "fmt"

func variabelEx() {

	//? variabel beserta tipe data

	a := "Hello a\n" //* penggunaan variabel shorthand

	var b string //* penggunaan variabel biasa

	b = "Hello b\n"

	var c = "Hello c\n" //* penggunaan variabel biasa

	var d string = "Hello d\n" //* penggunaan variabel biasa

	var ( //* multiple variabel
		firstName = "Rizqi"
		lastName  = "Qito\n"
	)

	const ( //* multiple constant
		satu  string = "satu"
		dua          = "dua"
		nilai        = 1000
	)

	//? alias type declaration
	type NoKTP string
	type Maried bool

	var noKTPku NoKTP = "6472052382392"
	var myMaried Maried = false

	fmt.Println(a, b, c, d, firstName, lastName, satu, dua, nilai, noKTPku, myMaried)
}

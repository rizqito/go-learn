package helper

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func HandsOnArray() {
	fmt.Println("----array----")
	var aa [2]string = [2]string{"Go", "C"} // array statis
	fmt.Println(aa[0])
	fmt.Println(aa[1])

	fmt.Println("----array-slices----")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// var s []int = primes[2:5] // get range data slice
	var s []int = primes[:] // get all data slice
	fmt.Println(s)

	fmt.Println("----array-slices-make----")
	// buat slice of int bernama a dengan panjang 5
	a := make([]int, 5)
	printSlice("a", a)
	// buat slice of int bernama b dengan panjang 0 kapasitas 5
	b := make([]int, 0, 5)
	printSlice("b", b)
	// buat variabel c dengan dua data awal b
	c := b[:2]
	printSlice("c", c)
	// buat variabel d dengan data b index ke 2 sampai 4
	d := c[2:5]
	printSlice("d", d)
}

package helper

import "fmt"

func HandsOnArray() {
	fmt.Println("----array----")
	var a [2]string = [2]string{"Go", "C"} // array statis
	fmt.Println(a[0])
	fmt.Println(a[1])

	fmt.Println("----array-slices----")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[3:5]
	fmt.Println(s)
}

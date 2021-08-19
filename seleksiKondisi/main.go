package main

import "fmt"

func main() {
	//* penggunaan if
	var point = 8
	if point == 8 {
		fmt.Println("point 8")
	} else if point == 7 {
		fmt.Println("point 7")
	} else {
		fmt.Println("Lain point")
	}

	//* penggunaan switch
	switch point {
	case 8:
		fmt.Println("Point Delapan")
	case 6, 7:
		fmt.Println("Point 6 dan 7")
	default:
		fmt.Println("lain point")
	}
}

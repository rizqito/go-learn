package helper

import (
	"fmt"
	"math"
)

func Sqrt() {
	x, y := float64(3), float64(4)
	c := math.Sqrt(x*x + y*y)
	z := uint(c)
	fmt.Println(x, y, z)
}

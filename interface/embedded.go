package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

//* embedded interface
type calc interface {
	hitung2d // call interface hitung2d
	hitung3d // call interface hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func embeddedTest() {
	var bangunRuang calc = &kubus{4}

	fmt.Println("==kubus==")
	fmt.Println("luas : ", bangunRuang.luas())
	fmt.Println("keliling : ", bangunRuang.keliling())
	fmt.Println("volume : ", bangunRuang.volume())
}

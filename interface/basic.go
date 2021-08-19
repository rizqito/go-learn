package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return (p.sisi + p.sisi) * 2
}

type lingkaran struct {
	jari float64
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jari, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * math.Pow(l.jari, 2)
}

func basic() {
	var bangunDatar hitung

	bangunDatar = persegi{5}
	fmt.Println("==persegi==")
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling : ", bangunDatar.keliling())

	bangunDatar = lingkaran{7}
	fmt.Println("==Lingkaran==")
	fmt.Println("jari-jari : ", bangunDatar.(lingkaran).jari) // memanggil nilai jari2 dari struct
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling : ", bangunDatar.keliling())
}

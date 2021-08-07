//interface = kumpulan method
package main

import "fmt"

// interface berguna untuk method bisa dipake berulang-ulang
type BangunDatar interface{
	Luas()
	Keliling()
}

// beberapa objek

// objek persegi
type Persegi struct{
	Sisi float64
}

// objek persegi panjang
type PersegiPanjang struct{
	Panjang float64
	Lebar float64
}

type Lingkaran struct{
	Jari2 float64
}

// method with receiver param = (biasanya digunakan untuk langsung inject ke struct)

// method persegi
func(p Persegi) Luas() float64{
	return p.Sisi * p.Sisi
}

func(p Persegi) Keliling() float64{
	return 4 * p.Sisi
}

// method persegi panjang
func(pp PersegiPanjang) Luas() float64{
	return pp.Panjang * pp.Lebar
}

func(pp PersegiPanjang) Keliling() float64{
	return 2 * (pp.Panjang + pp.Lebar)
}

// method lingkaran
func(l Lingkaran) Luas() float64{
	return 3.14 * (l.Jari2 * l.Jari2)
}

// function persegi
func hitungPersegi(p Persegi){
	fmt.Println("Keliling Persegi : ", p.Keliling())
	fmt.Println("LUas Persegi : ", p.Luas())
}

// function persegi panjang
func hitungPersegiPanjang(pp PersegiPanjang){
	fmt.Println("Keliling Persegi Panjang : ", pp.Keliling())
	fmt.Println("LUas Persegi Panjang : ", pp.Luas())
}

// function lingkaran
func hitungLingkaran(l Lingkaran){	
	fmt.Println("LUas Lingkaran : ", l.Luas())
}

func main(){
	fmt.Println("----------Persegi---------")
	hitungPersegi(Persegi{Sisi: 6})

	fmt.Println("----------Persegi Panjang---------")
	hitungPersegiPanjang(PersegiPanjang{Panjang: 6, Lebar: 5})

	fmt.Println("----------Lingkaran---------")
	hitungLingkaran(Lingkaran{Jari2: 6})
}
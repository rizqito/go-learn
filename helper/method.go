package helper

import "fmt"

// Luas(), Keliling() diamil dari interface interface.go

type Persegi1 struct {
	Sisi float64
}

type PersegiPanjang1 struct {
	Panjang float64
	Lebar   float64
}

type Lingkaran1 struct {
	r float64
}

// method with receiver param = (biasanya digunakan untuk langsung inject ke struct)
func (s Persegi1) Luas() float64 {
	return s.Sisi * s.Sisi
}

//method luas ersegi pajang
func (s PersegiPanjang1) Luas() float64 {
	return s.Panjang * s.Lebar
}

// method keliling persegi panjang
func (s PersegiPanjang1) Keliling() float64 {
	return 2 * (s.Panjang + s.Lebar)
}

//method luas lingkaran
func (s Lingkaran1) Luas() float64 {
	return 3.14 * (s.r * s.r)
}

func BelajarMethod() {
	fmt.Println("----------persegi---------")
	persegi := Persegi1{6}
	fmt.Println(persegi.Luas())

	// buar struct persegi panjang
	// isi struct nya dengan atribut persegi panjang
	// buat method with receiver untuk menghitung luas sama keliling
	// call method with receiver nya dengan print
	fmt.Println("----------persegi panjang----------")
	s := PersegiPanjang1{Panjang: 4, Lebar: 5}
	fmt.Println("Luas : ", s.Luas())
	fmt.Println("Keliling : ", s.Keliling())

	fmt.Println("----------lingkaran----------")
	r := Lingkaran1{r: 4}
	fmt.Println("Luas : ", r.Luas())
}

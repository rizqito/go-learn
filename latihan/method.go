package latihan

import "fmt"

type Persegi struct{
	Sisi float64
}

type PersegiPanjang struct{
	Panjang float64
	Lebar   float64
}

type Lingkaran struct{
	r float64
}

// method with receiver param = (biasanya digunakan untuk langsung inject ke struct)
func (s Persegi) luas() float64{
	return s.Sisi * s.Sisi
}

//method luas ersegi pajang
func (s PersegiPanjang) luas() float64{
	return s.Panjang * s.Lebar
}

// method keliling persegi panjang
func (s PersegiPanjang) keliling() float64{
	return 2 * (s.Panjang + s.Lebar)
}

//method luas lingkaran
func (s Lingkaran) luas() float64{
	return 3.14 * (s.r * s.r)
}

func main(){
	fmt.Println("----------persegi---------")
	persegi := Persegi{6}
	fmt.Println(persegi.luas())

	// buar struct persegi panjang
	// isi struct nya dengan atribut persegi panjang
	// buat method with receiver untuk menghitung luas sama keliling
	// call method with receiver nya dengan print
	fmt.Println("----------persegi panjang----------")
	s := PersegiPanjang{Panjang: 4, Lebar: 5}	
	fmt.Println("Luas : ", s.luas())
	fmt.Println("Keliling : ", s.keliling())

	fmt.Println("----------lingkaran----------")
	r := Lingkaran{r: 4}	
	fmt.Println("Luas : ", r.luas())
}
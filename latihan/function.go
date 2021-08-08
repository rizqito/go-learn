package latihan

import (
	"fmt"
	"math"
)

// ada 2 tipe method function golang

func showName(name string) string{ // #1 tipe method menggunakan return
	return name
}

func akarKuadrat(numberONe float64) float64{
	hasil := math.Sqrt(numberONe)

	return hasil
}

func calc(numberONe, numberTwo int) (int, error){ // jika parameter 1 tipe data dan 2 return
	// calcTambah := numberONe + numberTwo
	// calcKurang := numberONe - numberTwo
	// calcKali   := numberONe * numberTwo
	// calcBagi   := numberONe / numberTwo
	calcModulo := numberONe % numberTwo, nil

	return calcModulo, err;
}

func tambah(numberONe int, numberTwo int){ // #2 tipe method tanpa return
	fmt.Println("hasil tambah ",numberONe + numberTwo)
}

func kurang(numberONe int, numberTwo int){ // #2 tipe method tanpa return
	fmt.Println("hasil kurang ",numberONe - numberTwo)
}

func kali(numberONe int, numberTwo int){ // #2 tipe method tanpa return
	fmt.Println("hasil kali ",numberONe * numberTwo)
}

func bagi(numberONe int, numberTwo int){ // #2 tipe method tanpa return
	fmt.Println("hasil bagi ",numberONe / numberTwo)
}

func modulo(numberONe int, numberTwo int){ // #2 tipe method tanpa return
	fmt.Println("hasil modulo ",numberONe % numberTwo)
}


func main(){
	fmt.Println("----------tipe method return-----------")
	calcModulo, err := calc(1,3)

	if(err != null){

	}

	fmt.Println("----------tipe method tanpa return-----------")
	//buat function tambah, bagi, kali, kurang, modulo
	tambah(5,6)
	kurang(5,6)
	kali(5,6)
	bagi(5,6)
	modulo(6,5)

	fmt.Println("----------tipe method return akar kuadrat-----------")
	//buat function menghitung akar kuadrat
	fmt.Println(akarKuadrat(2))
}
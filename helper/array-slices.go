package helper

import "fmt"

func ArraySlices() {
	fmt.Println("----------array statis----------")
	var score [3]int = [3]int{40, 50, 60} // array statis

	fmt.Println(score) // output seperti biasa

	for i := 0; i < len(score); i++ { // get data with for
		fmt.Println(score[i])
	}

	for _, value := range score { // get data with foreach
		fmt.Println(value)
	}

	fmt.Println("----------array dinamis / slices----------")
	var count []int          //deklarasi variabel array dinamis / slices
	count = append(count, 1) //perintah untuk nambahin/isi value dalam array
	count = append(count, 2)
	count = append(count, 3)
	for i := 4; i < 10; i++ { //perintah untuk nambahin/isi value dalam array menggunakan looping
		count = append(count, i)
	}
	fmt.Println(count)

	// misalkan ada data buah dibawah ini
	// semangka, pisang, apel, belimbing
	// simpan data array tersebut dalam array dinamis
	// buat variabel penampung hasil looping
	// loop data array diatas menggunakan looping
	// print array setelah  di append di looping
	// print panjang dari array dinamis tersebut
	fmt.Println("----------latihan----------")

	buah, jml := []string{"semangka", "pisang", "apel", "belimbing"}, []string{}

	for _, value := range buah { // get data with foreach
		jml = append(jml, value)
	}

	fmt.Println(jml)
	fmt.Println("jumlah data ", len(jml))
}

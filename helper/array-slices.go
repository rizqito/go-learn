package helper

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

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

	fmt.Println("----array-slices----")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[3:5]
	fmt.Println(s)

	fmt.Println("----array-slices-make----")
	// buat slice of int bernama a dengan panjang 5
	a := make([]int, 5)
	printSlice("a", a)
	// buat slice of int bernama b dengan panjang 0 kapasitas 5
	b := make([]int, 0, 5)
	printSlice("b", b)
	// buat variabel c dengan dua data awal b
	c := b[:2]
	printSlice("c", c)
	// buat variabel d dengan data b index ke 2 sampai 4
	d := c[2:5]
	printSlice("d", d)

	fmt.Println("----array-slices-for-range----")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// var pow2 = []int{128, 64, 32, 16, 8, 4, 2, 1}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

package latihan

import (
	"fmt"
)

func main() {
	fmt.Println("macam-macam perintah")

	// tipe data variabel
	fmt.Println("-----tipe data variabel-----")
	var name string //string tipe data karakter
	var age int     //integer tipe data numerik
	var isPredicate bool //tipe data boolean
	name = "qito"
	age = 22
	isPredicate = false
	fmt.Println("Nama saya", name)
	fmt.Println("Umur saya", age)
	fmt.Println("Keterangan", isPredicate)

	fmt.Println("-----variabel shorthand-----")
	count := 20 //shorthand variabel (deklarasi tipe data otomatis)
	fmt.Println(count)

	fmt.Println("-----kondisi if-----")
	artist := "Muse" //if
	if artist == "Blackpink" {
		fmt.Println("Korea")
	} else {
		fmt.Println("not Korea")
	}

	fmt.Println("-----kondisi switch-----")
	switch artist { //switch
	case "Blackpink":
		fmt.Println("Korea")
		break
	default:
		fmt.Println("Not Korea")
	}

	fmt.Println("-----perulangan for-----")
	for i := 0; i < 10; i++ { //perulangan for
		fmt.Println(i)
	}

	fmt.Println("-----deklarasi array dengan for range-----")
	array := []string{"Apple", "Mango"} //deklarasi array

	for i, v := range array { //for range
		fmt.Println(i, v)
	}

	fmt.Println("-----pointer-----")
	number := 20
	pointerNumb := &number    //alokasi data ke memori
	fmt.Println(*pointerNumb) // akses nilai
	fmt.Println(pointerNumb)  // akses alamat memori
}

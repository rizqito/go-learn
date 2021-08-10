//map pada golang itu object

package helper

import "fmt"

func Map() {
	ages := make(map[string]int) // deklarasi variabel map menggunakan short hand

	// var ages map[string]int //deklarasi variabel biasa
	// ages = map[string]int{}

	ages["ani"] = 24
	ages["budi"] = 25
	ages["agung"] = 22

	fmt.Println(ages)

	for key, value := range ages { // foreach data
		fmt.Println("Key : ", key, "Value : ", value)
	}

	fmt.Println("Map")
	// buat variabel m dengan map key: string dan value: int
	m := make(map[string]int)
	// beri nilai "Answer" dengan 42
	m["Answer"] = 42
	fmt.Println("The value : ", m["Answer"])
	// beri nilai "Answer" dengan 48
	m["Answer"] = 48
	fmt.Println("The value : ", m["Answer"])
	// hapus "answer"
	delete(m, "Answer")
	fmt.Println("The value : ", m["Answer"])
	// gunakan pengecekan: elem, ok = m[key]
	v, ok := m["Answer"]
	fmt.Println("The value : ", v, "Present?", ok)
}

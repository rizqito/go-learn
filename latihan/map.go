//map pada golang itu object

package latihan

import "fmt"

func main(){
	ages := make(map[string]int)  // deklarasi variabel map menggunakan short hand

	// var ages map[string]int //deklarasi variabel biasa
	// ages = map[string]int{}

	ages["ani"] = 24
	ages["budi"] = 25
	ages["agung"] = 22
	
	fmt.Println(ages)

	for key, value := range ages{ // foreach data
		fmt.Println("Key : ", key, "Value : ", value)
	}
}
package helper

import (
	"fmt"
	"math/rand"
)

func Random() {
	rand.Seed(20)
	random := rand.Intn(100)
	fmt.Println("Nomor Favorit Saya Adalah : ", random)
}

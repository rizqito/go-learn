package main

import (
	"fmt"
	"math/rand"
)

func main(){	
	rand.Seed(20)
	random := rand.Intn(100)
	fmt.Println("Nomor Favorit Saya Adalah : ",random)	
}
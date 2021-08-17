package goRoutine

import "fmt"

func HandsOn3() {
	pesan := make(chan string, 2)
	pesan <- "Hello"
	pesan <- "goroutine"
	pesanAwal := <-pesan
	pesanAkhir := <-pesan
	fmt.Println(pesanAwal, pesanAkhir)
}

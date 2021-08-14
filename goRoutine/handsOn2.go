package goRoutine

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(234 + <-ch)
}

func HandsOn2() {
	fmt.Println("start Main Method")
	ch := make(chan int) // *membuat channel
	go myfunc(ch)        // *panggilan fungsi dengan goroutine
	ch <- 23             // *masukkan angka 23 ke dalam channel
	fmt.Println("End main method")
}

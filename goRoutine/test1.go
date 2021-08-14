package goRoutine

import (
	"fmt"
	"time"
)

func showName(name string) {
	fmt.Println("My name is : ", name)
}

func sendEmail(email string) {
	fmt.Println("EMail : ", email)
}

func Test() {
	// showName("Qito")
	// go showName("Rizqi")

	sendEmail("qito25@gmail.com")
	go sendEmail("rizqiqito@gmail.com")
	time.Sleep(2 * time.Second) //menampilkan  method call go routine
}

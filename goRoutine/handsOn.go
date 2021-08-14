package goRoutine

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func HandsOn() {
	go say("holla") // *panggil say("hello") dengan goroutine baru
	say("hallo")    // *panggil say("hello") seperti biasa
}

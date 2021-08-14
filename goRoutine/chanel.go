package goRoutine

import (
	"fmt"
	"time"
)

// chaneling
func Chanel() {
	channels := make(chan int, 2)

	go func() { // boleh pake boleh nggak
		channels <- 56
		channels <- 65
	}()

	result := <-channels
	result2 := <-channels

	fmt.Println(result)
	fmt.Println(result2)

	time.Sleep(1 * time.Second)
}

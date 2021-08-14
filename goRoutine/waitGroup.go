package goRoutine

import (
	"fmt"
	"sync"
	"time"
)

func tampilName(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("showName is starting..")

	fmt.Println("my name is : ", name)
	time.Sleep(time.Second)

	fmt.Println("Showname is done..")
}

func WaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go tampilName("Alex", &wg)
	}

	wg.Wait()
}

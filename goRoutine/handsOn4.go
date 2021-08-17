package goRoutine

import (
	"fmt"
	"sync"
	"time"
)

func showLastName(name string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Worker is starting..")
	fmt.Println("My name is : ", name)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Worker is done..")
}

func HandsOn4() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ { // Spawn goroutine sejumlah 10 item
		wg.Add(1)
		go showLastName("Alex", &wg) // Handle spawning goroutine using WaitGroup
	}

	wg.Wait()
}

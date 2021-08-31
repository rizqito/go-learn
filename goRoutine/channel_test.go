package goRoutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel dengan tipe data string
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// mengirim data ke channel
		channel <- "Rizqi Qito"
		fmt.Println("selesai kirim data ke channel")
	}()

	// menerima data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizqi Qito"
}

func TestChannelAsParameter(t *testing.T) {
	// membuat channel dengan tipe data string
	channel := make(chan string)
	defer close(channel)

	// manggil function dengan parameter channel
	go GiveMeResponse(channel)

	// menerima data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// function untuk hanya mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Qito Rizqi"
}

// function untuk hanya menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

// buffered channel = channel yang bisa ditampung sesuai dengan batas yang kita tentukan
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "qito"
		channel <- "rizqi"
		channel <- "rizqito"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

// range channel = cocok digunakan untuk test jumlah data yang tidak diketahui batasnya
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data ", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}

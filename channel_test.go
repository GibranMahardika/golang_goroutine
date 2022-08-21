package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// cara run, go test -v -run=NamaFunction

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// channel <- "Gibs"

	// data := <-channel

	// fmt.Println("data : ", data)
	// fmt.Println("channel : ", <-channel)
	defer close(channel) //menggunakan defer untuk force close channel

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Gibran Mahardika"
		fmt.Println("Selesai mengirim data ke channel")
	}() //membuat goroutine dengan anonymous function

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel sebagai Parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Gibran Mahardika"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //menggunakan defer untuk force close channel

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel In dan Out

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Gibran Mahardika"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(8 * time.Second)
}

//Buffered Channel

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Gibs"
	channel <- "Gibran"
	channel <- "Mahardika"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}

//Buffered Channel menggunakan Goroutine
func TestBufferedChannelGoroutine(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Gibs"
		channel <- "Gibran"
		channel <- "Mahardika"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Selesai")
}

//Range Channel

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

// Select CHannel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	select {
	case data := <-channel1:
		fmt.Println("Data dari channel1 ", data)
	case data := <-channel2:
		fmt.Println("Data dari channel2 ", data)
	}
	select {
	case data := <-channel1:
		fmt.Println("Data dari channel1 ", data)
	case data := <-channel2:
		fmt.Println("Data dari channel2 ", data)
	}
}

func TestSelectChannelLooping(t *testing.T) {
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
			fmt.Println("Data dari channel1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2 ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

// Default Select
func TestDefaultSelectChannel(t *testing.T) {
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
			fmt.Println("Data dari channel1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2 ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}

}

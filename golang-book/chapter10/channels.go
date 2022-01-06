package main

import (
	"fmt"
	"time"
)

func pinger(channel chan<- string) {
	for i := 0; ; i++ { // another way to create an infinite loop
		// `<-` operator used for sending and receiving messages on the channel
		channel <- "ping" // send "ping" to channel
	}
}

func ponger(channel chan string) {
	for i := 0; ; i++ {
		channel <- "pong"
	}
}

func printer(channel <-chan string) {
	for {
		message := <-channel // receive message from channel and store it in "message"
		// the receiver must be ready first for receiveing the message before a channel can send it (blocking)
		fmt.Println(message)
		// fmt.Println(<-channel) => another way, receive and print directly
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// channels provide a way for two goroutines to communicate with each others and synchronize their execution
	// using the keyword `channel` followed by a type (string in this example)
	channel := make(chan string)
	// cheatsheets
	// `chan` => send and receive message (bi-directional)
	// `chan<-` => only send message
	// `<-chan` => only receive message

	go pinger(channel)
	go ponger(channel)
	go printer(channel)

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "from channel 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			channel2 <- "from channel 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// switch-like statement, but for channel
			// it will pick the first-ready channel (send/receive) and will block until one is ready
			// exception if we have default case
			select {
			case mesaage1 := <-channel1:
				fmt.Println(mesaage1)
			case message2 := <-channel2:
				fmt.Println(message2)
			case <-time.After(time.Second): // select often used to implement timeout (no channel is ready case)
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	// normal channel => synchronous (both sides will wait)
	// buffered channel => asynchronous (will not wait, unless the channel is already full)
	// buffered channel with capacity of 1
	// bufferedChannel := make(chan string, 1)

	var input string
	fmt.Scanln(&input)
}

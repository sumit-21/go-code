package main

import (
	"fmt"
	"time"
)

func dataProducer(myChannel chan string) {
	myChannel <- "Test Data"
}

func dataConsumer(myChannel chan string) {
	value := <-myChannel
	fmt.Print("Data received by consumer", value)
}

func testChannel() {
	myChannel := make(chan string)
	for i := 0; i < 10; i++ {
		go dataProducer(myChannel)
	}
	for i := 0; i < 6; i++ {
		dataConsumer(myChannel)
		fmt.Println(", i =", i)
	}
	fmt.Println("Data received by main", <-myChannel)
}

func testBufferedChannel() {
	myBufferedChannel := make(chan string, 2)
	// send data
	myBufferedChannel <- "Send"
	myBufferedChannel <- "My"
	myBufferedChannel <- "Data"
	// receive data
	fmt.Println(<-myBufferedChannel + " " + <-myBufferedChannel + " " + <-myBufferedChannel)
}

func testChannelSynchronization() {
	myChannel := make(chan bool)

	go func(tempChannel chan bool) {
		fmt.Println("Starting routine")
		time.Sleep(3 * time.Second)
		fmt.Println("Closing routine")
		tempChannel <- true
	}(myChannel)

	// wait for above routine to finish
	<-myChannel
	fmt.Println("Main finished")
}

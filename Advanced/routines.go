package main

import (
	"fmt"
	"time"
)

func printData(data string) {
	for i := 0; i < 5; i++ {
		fmt.Println(data, i+1)
	}
}

func testRoutines() {
	// sync call
	printData("Sync call")
	// go routine
	go printData("Go routine call")
	// go routine with anonymous function call
	go func(data string) {
		for i := 0; i < 5; i++ {
			fmt.Println(data, i+1)
		}
	}("Go routine anonymous call")
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

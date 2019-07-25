package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// <name> chan <datatype>

	// send in a goroutine
	go func() {
		c <- 1
	}()

	// sniff
	val := <-c
	fmt.Println(val)

	// send in a goroutine
	go func() {
		c <- 2
	}()

	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val)

	fmt.Println(c)
}

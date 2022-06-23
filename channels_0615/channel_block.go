package main

import (
	"fmt"
	"time"
)

func Block(){
	ch1 := make(chan int,10)
	go sendD(ch1)
	for {
		fmt.Println(<-ch1)
	}
}

func sendD(ch chan int)  {
	for i := 1; i < 15; i++ {
		ch <- i
		time.Sleep(1e9)
	}
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
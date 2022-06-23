package main

import (
	"fmt"
)

func MessageT()  {
	ch := make(chan string)
	//go sendData(ch)
	go getData(ch)
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
	ch <- "7"
}

func sendData(ch chan string)  {
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
	ch <- "7"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf(input)
	}
}
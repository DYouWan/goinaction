package main

func SemaphoreMain(ch chan int,array ...int) {
	var count int
	for i := 0; i < len(array); i++ {
		count += array[i]
	}
	ch <- count
}

func semaphore(ch chan int,array ...int) {
	var count int
	for i := 0; i < len(array); i++ {
		count += array[i]
	}
	ch <- count
}
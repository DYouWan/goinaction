package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s2 := make([]int,1,10)

	sw := Magnify(10, s1)
	fmt.Println(sw)
}

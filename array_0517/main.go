package main

import (
	"fmt"
)

func main() {
	//1.求数组[1, 3, 5, 7, 8]所有元素的和
	//2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	sum := 0
	array := []int{1, 3, 5, 7, 8}
	for _, v := range array {
		sum += v
	}
	fmt.Println("元素的和为: ", sum)

	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array); j++ {
			if array[i]+array[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}

package main

import "fmt"

func main() {
	var (
		i, j = 1,1
	)
	for {
		if i == j {
			fmt.Printf("%d*%d=%d\n", i, j, i*j)
			if i < 9 && j < 9 {
				j++
				i = 1
			} else {
				break
			}
		} else {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
			i++
		}
	}

	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d  ", x, y, x*y)

		}
		fmt.Println()
	}

}

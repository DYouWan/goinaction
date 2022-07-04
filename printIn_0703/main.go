package main

import (
	"fmt"
	"os"
)

func main()  {
	println("123")
}
func Println(a ...interface{}) (n int, err error) {
	println("draven")
	return fmt.Fprintln(os.Stdout, a...)
}

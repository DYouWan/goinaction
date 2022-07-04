package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('S')
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(val, "\n")
	fmt.Println(len(lines))
}

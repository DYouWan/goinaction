package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num      = 10000000  // 测试数组的长度
	rangeNum = 100000 // 数组元素大小范围
)

func main() {
	arr := GenerateRand()
	start(arr)
}

func start(arr [] int) {
	defer timeCost(time.Now())
	Bubble(arr)
}

func timeCost(start time.Time) {
	tc := time.Since(start)
	fmt.Printf("time cost = %v\n", tc)
}

// GenerateRand 生成随机数组
func GenerateRand() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = randSeed.Intn(rangeNum)
	}
	return arr
}

func Bubble(arr []int) {
	size := len(arr)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}
}


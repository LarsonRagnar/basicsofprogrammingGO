package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [100]int
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range arr {
		value := rnd.Intn(20)
		arr[i] = value
	}
	fmt.Println(arr)
	idx := arr[0]
	repeatNum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == idx {
			repeatNum++
			//		for j := i + 1; j < len(arr); j++ {
			//			if arr[i] == arr[j] {
			//				repeatNum++
			//			}
		}
	}
	fmt.Println(repeatNum)
	sliceArr := make([]int, repeatNum)
	fmt.Println(sliceArr)
}

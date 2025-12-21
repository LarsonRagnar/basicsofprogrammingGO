package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int)
	// go func() {
	// 		<-strChan
	// }()
	// strChan <-""
	// close(strChan)
	// str,ok := <- strChan
	go func() {
		for i := 0; i < 4; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	// for {
	// 	val, ok := <-intChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(val)
	// }
	for val := range intChan {
		fmt.Println(val)
	}
}

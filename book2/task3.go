package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Testfor()
}

func Testfor() {
	slice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		rndNum := rand.Intn(100)
		slice = append(slice, rndNum)
	}
	fmt.Println(slice)
	for _, v := range slice {
		switch {
		case v%2 == 0:
			fmt.Println("Два!")
		case v%3 == 0:
			fmt.Println("Три!")
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Шесть!")
		default:
			fmt.Println("Неважно!")
		}
	}
}

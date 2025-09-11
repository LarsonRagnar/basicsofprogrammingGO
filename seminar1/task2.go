package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func displayPrimes(n int) {
	if isPrime(n) {
		fmt.Println("Число ", n, " является простым")
	} else {
		fmt.Println("Число ", n, " не является простым")
	}
}

func main() {
	fmt.Print("Введите число: ")
	var n int
	if _, err := fmt.Fscan(os.Stdin, &n); err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}
	displayPrimes(n)
}

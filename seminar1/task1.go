package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("not an integer: %v", err)
	}
	if isPrime(n) {
		fmt.Println("prime")
	} else {
		fmt.Println("not prime")
	}

	// Pause so the console window stays open on double-click
	fmt.Println("Press Enter to exit...")
	reader.ReadString('\n')
}

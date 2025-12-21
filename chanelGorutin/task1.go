package main

import (
	"fmt"
)

func main() {
	respChan := make(chan string)
	go putBook(respChan)
	go deliverBook(respChan)
	burnBook()
}

func putBook(rchan chan string) {
	fmt.Println("складываю книги")
}

func deliverBook(rchan chan string) {
	fmt.Println("доставляю книги")
}

func burnBook() {
	fmt.Println("Сжигаю книги")
}

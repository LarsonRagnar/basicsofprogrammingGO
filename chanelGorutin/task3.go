package main

import (
	"fmt"
)

func main() {
	fc := putBook()
	sc := deliveryBook(fc)
	tc := burnBook(sc)
	fmt.Println(<-tc)

}

func putBook() chan string {
	firstChan := make(chan string)
	go func() {
		firstChan <- "Собираю книги"
	}()
	return firstChan
}

func deliveryBook(firstChan chan string) chan string {
	fmt.Println(<-firstChan)
	secondChan := make(chan string)
	go func() {
		secondChan <- "везу тележку"
	}()
	return secondChan
}

func burnBook(secondChan chan string) chan string {
	fmt.Println(<-secondChan)
	thirdChan := make(chan string)
	go func() {
		thirdChan <- "сжигаю книги"
	}()
	return thirdChan
}

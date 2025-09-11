package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	_ "unicode"
)

func main() {
	fmt.Println("Введите предложение: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}
	var wordLine, symbolNumber, bytesNumber int
	bytesText := []byte(text)
	bytesNumber = len(bytesText)
	words := strings.Fields(text)
	wordLine = len(words)
	for _, symbol := range words {
		symbolNumber += len([]rune(symbol))
	}
	fmt.Println("Количество символов: ", symbolNumber)
	fmt.Println("Количество байт: ", bytesNumber)
	fmt.Println("Количество слов: ", wordLine)
}

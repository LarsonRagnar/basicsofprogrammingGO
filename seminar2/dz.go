package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Вы не ввели предложение")
	}

	// Создаем map для слов
	textMap := make(map[int]string)
	for idx, arg := range os.Args[1:] {
		textMap[idx] = strings.ToLower(arg)
	}
	fmt.Println("Исходные слова:", textMap)

	// Создаем map для подсчета частоты букв
	letterCount := make(map[rune]int)
	totalLetters := 0

	// Подсчитываем частоту каждой буквы
	for _, word := range textMap {
		for _, symbol := range word {
			// Проверяем, что это буква (не символ, не знак препинания)
			if unicode.IsLetter(symbol) {
				letterCount[symbol]++
				totalLetters++
			}
		}
	}
	fmt.Println("map с буквами и их количеством с типом rune:", letterCount)
	fmt.Println("Колличество букв:", totalLetters)
	for letter, count := range letterCount {
		fmt.Printf("'%c': %d раз\n", letter, count)
	}

	// Создаем map для процентов (ключ - буква, значение - процент)
	percentageMap := make(map[rune]float64)

	// Вычисляем проценты для каждой буквы
	for letter, count := range letterCount {
		percentage := float64(count) / float64(totalLetters) * 100
		percentageMap[letter] = percentage
	}

	fmt.Println("\nMap с процентами (ключ - буква, значение - процент):")
	for letter, percentage := range percentageMap {
		fmt.Printf("'%c' -> %.2f%%\n", letter, percentage)
	}
}

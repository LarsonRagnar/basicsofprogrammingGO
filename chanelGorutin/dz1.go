// Задание 1. Конвейер
// Цели задания:
// -Научиться работать с каналами и горутинами.
// -Понять, как должно происходить общение между потоками.
// Что нужно сделать:
// -Реализуйте паттерн-конвейер:
// Программа принимает числа из стандартного ввода в бесконечном цикле
// и передаёт число в горутину.
// -Квадрат: горутина высчитывает квадрат этого числа
// и передаёт в следующую горутину.
// -Произведение: следующая горутина умножает квадрат числа на 2.
// -При вводе «стоп» выполнение программы останавливается.
// Советы и рекомендации:
// Воспользуйтесь небуферизированными каналами и waitgroup.
// Что оценивается:
// Ввод : 3
// Квадрат : 9
// Произведение : 18

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Введите число: ")
	reader := bufio.NewReader(os.Stdin)
	inputChan := make(chan int)
	squareChan := make(chan int)
	// doubleChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range inputChan {
			square := num * num
			fmt.Printf("Квадрат: %d\n", square)
			squareChan <- square
		}
		close(squareChan)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for square := range squareChan {
			result := square * 2
			fmt.Printf("Произведение: %d\n", result)
			// doubleChan <- result
		}
	}()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for result := range doubleChan {
	// 		fmt.Printf("Итоговый результат: %d\n", result)
	// 	}
	// }()

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "стоп" {
			fmt.Println("Программа завершена")
			close(inputChan) // Закрываем входной канал
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число или 'стоп' для выхода")
			continue
		}

		// Отправляем число в конвейер
		fmt.Printf("Ввод: %d\n", num)
		inputChan <- num
	}

	// Ждем завершения всех горутин
	wg.Wait()
}

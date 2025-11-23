package main

// *Задание 1. Работа с файлами**
// Что нужно сделать:
// -Напишите программу, которая на вход получала бы строку,
// введённую пользователем,
// а в файл писала № строки, дату и сообщение в формате:
// 2020-02-10 15:00:00 продам гараж.
// -При вводе слова exit программа завершает работу.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	f, err := os.Create("log1.log")
	if err != nil {
		fmt.Println("Неудалось создать файл", err)
		return
	}
	defer f.Close()
	fmt.Println("Введите текст:\n Для завершения введите exit")
	f.WriteString("Введите текст:\n Для завершения введите exit\n")
	const timeFormat = "2006-01-02 15:04:05"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			fmt.Println("Завершение работы")
			break
		}
		time.Sleep(1)
		timestamp := time.Now().Format(timeFormat)
		logEntry := fmt.Sprintf("[%s] %s\n", timestamp, line)
		if _, err := f.WriteString(logEntry); err != nil {
			fmt.Println("Ошибка записи в файл: ", err)
			return
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("ОШибка чтения ввода", err)
	}
	readFile()
	createFileOnlyRead("log.log")
}

// -Напишите программу, которая читает и выводит в консоль строки
// из файла, созданного в предыдущей практике,
// без использования ioutil.
// Если файл отсутствует или пуст,
// выведите в консоль соответствующее сообщение.
// Рекомендация:
// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

func readFile() {
	r, err := os.Open("log1.log")
	if err != nil {
		fmt.Println("НЕудалось открыть файл")
		return
	}
	defer r.Close()
	fileInfo, err := r.Stat()
	if err != nil {
		fmt.Printf("Не удалось получить информацию о файле: %v\n", err)
		return
	}

	fileSize := fileInfo.Size()
	fmt.Printf("Размер файла '%s': %d байт\n", fileInfo.Name(), fileSize)
	buf := make([]byte, fileSize)
	_, err = r.Read(buf)
	if err != nil {
		fmt.Println("Не смог прочитать последовательность байтов из файла", err)
		return
	}
	fmt.Println(string(buf))

}
func createFileOnlyRead(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Не удалос создать файл", err)
		return
	}
	err = os.Chmod(name, 0444)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer.WriteString("Попытка записи")

}

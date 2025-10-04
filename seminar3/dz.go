package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Url struct {
	name string
	url  string
	tags string
	date string
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")
	urlList := make(map[string]Url)
	t := time.Now()
	t.Format(time.DateTime)

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}
			urlList[args[0]] = Url{
				name: args[1],
				url:  args[0],
				tags: args[2],
				date: t.Format("2006-01-02 15:04:05"),
			}

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			for _, url := range urlList {
				fmt.Println(url.name)
				fmt.Println(url.url)
				fmt.Println(url.tags)
				fmt.Println(url.date)
			}
			fmt.Println(len(urlList))
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			// ключом карты является URL; очищаем ввод от пробелов и перевода строки
			key := strings.TrimSpace(text)
			if _, ok := urlList[key]; ok {
				delete(urlList, key)
				fmt.Println("Удалено:", key)
			} else {
				fmt.Println("Не найдено:", key)
			}
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}

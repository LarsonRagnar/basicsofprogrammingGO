package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("ВВедите число от 1 до 100")
	b.WriteString("Введите число от 1 до 100 \n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			b.WriteString(fmt.Sprintf("ВВедено число %d", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
				b.WriteString("Число должно быть в диапозоне от 1 до 100 \n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Супер введено правильно число")
			b.WriteString("Супер введено правильно число \n")
			break
		} else if answer < n {
			fmt.Println("Загаданое число больше")
			b.WriteString("Загаданое число больше \n")
		} else if answer > n {
			fmt.Println("Загаданое число меньше")
			b.WriteString("Загаданое число меньше \n")
		}
	}
	filename := "log.txt"
	if err := ioutil.WriteFile(filename, b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохраненый лог: ")
	fmt.Println(string(resultBytes))
}

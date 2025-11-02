package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("не смог создать файл сорян ", err)
		return
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("ВВедите число от 1 до 100")
	file.WriteString("Введите число от 1 до 100 \n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			file.WriteString(fmt.Sprintf("ВВедено число %d", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
				file.WriteString("Число должно быть в диапозоне от 1 до 100 \n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Супер введено правильно число")
			file.WriteString("Супер введено правильно число \n")
			return
		} else if answer < n {
			fmt.Println("Загаданое число больше")
			file.WriteString("Загаданое число больше \n")
		} else if answer > n {
			fmt.Println("Загаданое число меньше")
			file.WriteString("Загаданое число меньше \n")
		}
	}

}

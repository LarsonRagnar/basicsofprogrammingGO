package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello"
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("НЕ смогли создать файл ", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
}

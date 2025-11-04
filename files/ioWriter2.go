package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var (
		username string
		password string
		age      int
	)
	fmt.Println("ВВедите име пользователя: ")
	fmt.Scan(&username)
	fmt.Println("ВВедите пароль: ")
	fmt.Scan(&password)
	fmt.Println("ВВедите возраст: ")
	fmt.Scan(&age)
	file, err := os.Create("credentian.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Ваш логин: %s", username))
	b.WriteString(fmt.Sprintf("Ваш пароль: %s", password))
	b.WriteString(fmt.Sprintf("Ваш возраст: %d", age))
	_, err = file.Write(b.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}

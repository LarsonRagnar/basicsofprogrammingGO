package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}
	filePath := os.Args[1]
	var fileName, fileExt string

	// Найти индексы
	lastSlashIndex := strings.LastIndex(filePath, "/")
	lastDotIndex := strings.LastIndex(filePath, ".")

	// Извлечь имя файла (все после последнего "/")
	if lastSlashIndex >= 0 {
		fileName = filePath[lastSlashIndex+1:]
	} else {
		fileName = filePath // если "/" не найден
	}

	// Извлечь расширение (все после последнего ".")
	if lastDotIndex >= 0 {
		fileExt = filePath[lastDotIndex:]
	} else {
		fileExt = "" // если "." не найден
	}

	fmt.Printf("имя файла: %s\n", fileName)
	fmt.Printf("расширение файла: %s\n", fileExt)
}

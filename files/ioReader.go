package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Не смог открыть открыть файл", err)
		return
	}
	defer f.Close()
	// buf := make([]byte, 256)
	// if _, err := io.ReadFull(f, buf); err != nil {
	// 	fmt.Println("Не смог прочитать последовательность байтов из файла", err)
	// 	return
	// }
	// fmt.Printf("%s\n", buf)
	buf := make([]byte, 128)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Не смог прочитать последовательность байтов из файла", err)
		return
	}
	fmt.Println(string(buf))

}

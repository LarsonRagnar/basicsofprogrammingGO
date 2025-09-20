package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo()
	echo1()
	echo2()
}

func echo() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = sep + os.Args[i]
		sep = " "
		fmt.Printf("%d %s\n", i, s)

	}
	elapsed := time.Since(start)
	fmt.Printf("echo() выполнилась за: %v\n\n", elapsed)
	//fmt.Println(s)
}

func echo1() {
	start := time.Now()
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s = sep + arg
		sep = " "
		fmt.Printf("%d %s\n", i, s)
	}
	elapsed := time.Since(start)
	fmt.Printf("echo1() выполнилась за: %v\n\n", elapsed)
	//fmt.Println(s)
}

func echo2() {
	start := time.Now()
	fmt.Printf("%s\n", strings.Join(os.Args[0:], " "))
	elapsed := time.Since(start)
	fmt.Printf("echo2() выполнилась за: %v\n\n", elapsed)
}

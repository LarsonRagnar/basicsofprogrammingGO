package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server is runnung")

	con, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(con)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("line recieved: ", line)
		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatal(err)
		}
	}
}

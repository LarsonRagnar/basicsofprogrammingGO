package main

import "fmt"

func main() {
	switch1()
	switch2()
	switch3()
	switch4()
}

func switch1() {
	words := []string{"a", "cow",
		"smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right lenght:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

}
func switch2() {
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 not 2")
		case 7:
			fmt.Println("exit the loop")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

func switch3() {
	words := []string{"popa", "dependence"}
	for _, word := range words {
		switch wordlen := len(word); {
		case wordlen < 5:
			fmt.Println(word, "is a short word!")
		case wordlen > 9:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is a medium word!")
		}
	}

}

func switch4() {
	for i := 0; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

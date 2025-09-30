package main

import "fmt"

func main() {
	task1()
	task2()
	task3()
}

func task1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привет"}
	slices1 := greetings[:2]
	slices2 := greetings[1:4]
	slices3 := greetings[3:]

	fmt.Printf("slices: %v\n %v\n %v\n", slices1, slices2, slices3)

}

func task2() {
	var message string
	message = "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Printf("runes: %v\n", string(runes[3]))
}

func task3() {
	type Employee struct {
		firstname string
		lastname  string
		id        int
	}
	batrak := Employee{
		firstname: "Ivan", lastname: "Ivanov", id: 1,
	}
	batrak2 := Employee{"Pety", "Petrov", 2}
	var batrak3 Employee
	batrak3.firstname = "Vasya"
	batrak3.lastname = "Vasin"
	batrak3.id = 3
	fmt.Printf("batrak: %v\n batrak2: %v\n batrak3: %v\n", batrak, batrak2, batrak3)
}

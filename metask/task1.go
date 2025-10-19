// создай перменные всех типов которые знаешь запони и выведи на экран
package main

import "fmt"

// Интерфейс
type Person interface {
	GetFullName() string
	IsAdult() bool
	DisplayInfo()
	ChangeAge(newAge int)
}

// Создать структиру и запонить
type children struct {
	firstName string
	lastName  string
	age       int
}

// Методы для структуры children
func (c children) GetFullName() string {
	return c.firstName + " " + c.lastName
}

func (c children) IsAdult() bool {
	return c.age >= 18
}

func (c children) DisplayInfo() {
	fmt.Printf("Ребенок: %s %s, Возраст: %d\n", c.firstName, c.lastName, c.age)
}

// функция изменяющая возраст с указателем
func (c *children) ChangeAge(newAge int) {
	c.age = newAge
}

// Функция, работающая с интерфейсом
func ProcessPerson(p Person) {
	fmt.Println("Обрабатываем:", p.GetFullName())
	if p.IsAdult() {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}
	p.DisplayInfo()

}

func main() {
	var (
		a int
		b float64
		c byte
		d bool
		e string
	)
	fmt.Printf("a=%d\n b=%f\n c=%v\n d=%t\n e=%s\n", a, b, c, d, e)
	a = 10
	b = 3.14
	c = 64
	d = true
	e = "fUck"
	fmt.Printf("a=%d\n,b=%f\n,c=%v\n,d=%t\n,e=%s\n", a, b, c, d, e)
	arr := [3]int{1, 2, 3}
	slice := arr[:]
	fmt.Printf("arr=%d\n slice=%d\n", arr, slice)
	slice = append(slice, 4, 5, 6, 7, 8)
	fmt.Println(slice)
	//приведи слайс к дрегому типу
	sliceFloat := make([]float64, len(slice))
	for i, v := range slice {
		sliceFloat[i] = float64(v)
	}
	fmt.Println(sliceFloat)
	//со3дать map и заполнить
	map1 := make(map[string]int)
	for j := 0; j < 10; j++ {
		key := fmt.Sprintf("item%d", j)
		map1[key] = j
	}
	fmt.Print(map1)
	//Простое добавление элементов
	map4 := make(map[string]int)
	map4["apple"] = 5
	map4["banana"] = 3
	map4["orange"] = 8
	// Способ 1: с помощью make()
	map5 := make(map[string]int)

	// Способ 2: литерал (сразу с данными)
	map2 := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	// Способ 3: пустой литерал
	map3 := map[string]int{}
	fmt.Println(map5, map2, map3)
	// Создать структиру и запонит

	ScollTwo := children{firstName: "lesha", lastName: "ivanov", age: 16}
	fmt.Println(ScollTwo)
	// Разные способы создания структур (как указатели)
	child1 := &children{firstName: "lesha", lastName: "ivanov", age: 16}
	child2 := &children{"anna", "petrova", 15}
	child3 := &children{firstName: "mike", age: 17}

	// Использование методов
	fmt.Println("Полное имя:", child1.GetFullName())
	fmt.Println("Совершеннолетний:", child1.IsAdult())
	child1.DisplayInfo()

	// Работа с интерфейсом
	fmt.Println("\n--- Работа с интерфейсом ---")
	ProcessPerson(child1)
	ProcessPerson(child2)
	ProcessPerson(child3)

	// Создание слайса интерфейсов
	people := []Person{child1, child2, child3}

	fmt.Println("\n--- Все люди ---")
	for i, person := range people {
		fmt.Printf("%d. ", i+1)
		ProcessPerson(person)
	}
	fmt.Println("\n--- Изменение возраста ---")
	child1.ChangeAge(19)
	child1.DisplayInfo()

}

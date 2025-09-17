// итерация по массиву с использованием range
/* переменные idx, value можно пропускать
for _, value := range arr {} // только значения
for idx := range arr {}      // только индексы
*/
package main

import (
	"fmt"
	"slices"
)

// итерируемся по массиву (срезу) получая значения и индекс
func main() {
	// объявляем и инициализируем срез строк
	arr := []string{"first", "second", "third"}

	// range отдаёт два значения: индекс (idx) и значение (value)
	for idx, value := range arr {
		// %d — целое число (индекс), %s — строка (значение)
		// \n — перевод строки, чтобы каждая пара печаталась с новой строки
		fmt.Printf("idx: %d, value: %s\n", idx, value)
	}

	// итерируем счётчик i и последовательно получаем доступ к элементам по индексу
	for i := 0; i < len(arr); i++ {
		fmt.Printf("idx: %d, value: %s\n", i, arr[i])
	}
	fmt.Println("--------------------------------")
	slice()
	fmt.Println("--------------------------------")
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println("--------------------------------")
	appendSlice()
	fmt.Println("--------------------------------")
	slicesTest()
	fmt.Println("--------------------------------")
	mapTest()
	fmt.Println("--------------------------------")
	nestedMap()
	fmt.Println("--------------------------------")
	mapOperations()
	fmt.Println("--------------------------------")
}

// срезы
func slice() {
	var slice []int
	var slice1 []int = []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	var slice3 = make([]int, 10)
	fmt.Println(slice, slice1, slice2, slice3)
}
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func appendSlice() {
	s := make([]int, 3, 6)
	fmt.Println(len(s), cap(s))
	s = append(s, 4, 5, 6)
	s[3] = 3
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}

// slices
func slicesTest() {
	s := []int{
		1, 2, 3, 4, 5,
	}
	s1 := slices.Delete(s, 2, 5)
	fmt.Println(s1)
	d := []int{1, 2, 3, 4, 5}
	d1 := slices.Insert(d, 2, 0, 0)
	fmt.Println(d1)
	type myType string
	v := []myType{"myFirst", "mySecond"}
	v = slices.Insert(v, len(v), "myThird")
	fmt.Println(v)
}

// map или отображения
func mapTest() {
	var m map[int]string
	m1 := make(map[int]string)
	m2 := map[int]string{1: "hello", 2: "world"}
	fmt.Println(m, m1, m2)
}

// вложенные map
func nestedMap() {
	nestedMap := make(map[string]map[string]int)
	nestedMap["first"] = make(map[string]int)
	nestedMap["first"]["a"] = 1
	nestedMap["first"]["b"] = 2
	nestedMap["second"] = make(map[string]int)
	nestedMap["second"]["x"] = 10
	nestedMap["second"]["y"] = 20
	fmt.Println(nestedMap)
}

// операции с map
func mapOperations() {
	myMap := make(map[string]int)
	myMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	values := make([]int, 0, len(myMap))
	keys := make([]string, 0, len(myMap))
	for k, v := range myMap {
		values = append(values, v)
		keys = append(keys, k)
	}
	fmt.Println(values, keys)
	slices.Sort(values)
	slices.Sort(keys)
	fmt.Println(values, keys)
}

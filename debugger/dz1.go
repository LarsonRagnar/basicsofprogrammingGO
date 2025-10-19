package main

import "fmt"

var arr [5]int
var arr1 [4]int

func sort(k []int) []int {
	n := len(k)
	for i := 0; i < n-1; i++ {
		// Находим индекс минимального элемента в оставшейся части массива
		minIndex := i
		for j := i + 1; j < n; j++ {
			if k[j] < k[minIndex] {
				minIndex = j
			}
		}
		// Меняем минимальный элемент с первым элементом несортированной части
		k[i], k[minIndex] = k[minIndex], k[i]
	}
	return k
}

func main() {
	arr = [5]int{5, 2, 8, 1, 9}
	arr1 = [4]int{8, 5, 6, 4}
	fmt.Println("До сортировки:", arr)
	slice1 := sort(arr[:]) // Преобразуем массив в слайс
	slice2 := sort(arr1[:])
	fmt.Println("После сортировки:", arr)
	newArr := make([]int, len(slice1)+len(slice2))
	copy(newArr, slice1)
	copy(newArr[len(slice1):], slice2)

	fmt.Println(sort(newArr))
}

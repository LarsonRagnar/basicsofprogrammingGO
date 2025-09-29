package main

import (
	"fmt"
)

type Order struct {
	Quantity     int
	Price        int
	CustomerName string
	ProductName  string
}

// func (r Order) PrintOrder() {
// 	total := r.Quantity * r.Price
// 	fmt.Printf("Заказ от %s: %dx %s (Итого: $%d)\n", r.CustomerName, r.Quantity, r.ProductName, total)
// }

func printOrder(r Order) {
	total := r.Quantity * r.Price
	fmt.Printf("Заказ от %s: %dx %s (Итого: $%d)\n", r.CustomerName, r.Quantity, r.ProductName, total)
}

func main() {
	printOrder(Order{Quantity: 12, Price: 300, CustomerName: "Иван", ProductName: "Книга"})
}

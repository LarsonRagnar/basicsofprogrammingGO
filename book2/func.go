package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	result, remainder, err := divi(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 4, 5, 6, 7, 8, 9))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 4, 5, 6, 7, 8}...))
	fmt.Println("---------------------------------------------------------------------------")
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expexpression := range expressions {
		if len(expexpression) != 3 {
			fmt.Println("invalid expression:", expexpression)
			continue
		}
		p1, err := strconv.Atoi(expexpression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expexpression[1]
		opFunc, ok := OpMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expexpression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
	fmt.Println("---------------------------------------------------------------------------")
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")

	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}

func divi(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return 0, 0, errors.New("connot divide by zero")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

var (
	add = func(i, j int) int { return i + j }
	sub = func(i, j int) int { return i - j }
	mul = func(i, j int) int { return i * j }
	div = func(i, j int) int { return i / j }
)

var OpMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

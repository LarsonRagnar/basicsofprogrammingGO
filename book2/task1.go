package main

import (
	"fmt"
	"math"
)

func main() {
	i := 20
	f := float64(i)
	fmt.Printf("f = %f, i = %d\n", f, i)
	fmt.Println("--------------------------------")
	constTest()
	fmt.Println("--------------------------------")
	thirdEnv()
}

func constTest() {
	const c = 20
	i := int(c)
	f := float64(c)
	fmt.Printf("i=%d, f=%f\n", i, f)
}

func thirdEnv() {
	var b byte
	var small int32
	var bigI uint64
	//	b = 127
	//	small = 2147483647
	//	bigI = 18446744073709551615
	b = math.MaxUint8
	small = math.MaxInt32
	bigI = math.MaxUint64
	b += 1
	small += 1
	bigI += 1
	fmt.Printf("b =%d, small=%d, bigI=%d\n", b, small, bigI)
}

package main

import (
	"fmt"
)

func modifyV(a int) int {
	a = a + 9
	return a * 2
}

func main() {
	// slice 示例
	originalSlice := []int{1, 2, 3}
	newSlice := originalSlice
	newSlice[0] = 10
	fmt.Println("Before modifying slice:", originalSlice)

	var a int = 2
	var po *int
	po = &a
	var b int = *po
	b += 2
	fmt.Println(modifyV(a))
	fmt.Println(a)
	fmt.Println(b)
}

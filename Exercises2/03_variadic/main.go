package main

import "fmt"

func main() {
	n := high(54, 32, 78, 21, 11)
	fmt.Println(n)
}

func high(x ...int) int {
	fmt.Println(x)
	var highest int
	for _, v := range x {
		if v > highest {
			highest = v
		}
	}
	return highest
}

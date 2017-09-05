package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	fmt.Println(divide(number))
}

func divide(x int) (int, bool) {
	return x / 2, even(x)
}

func even(z int) bool {
	if z%2 == 0 {
		return true
	}
	return false
}

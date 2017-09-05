package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	h, even := half(number)
	fmt.Println(h, even)
}

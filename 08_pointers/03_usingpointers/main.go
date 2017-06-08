package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%d \n", &a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(a)
}

package main

import "fmt"

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	sum := 0
	f := fibo()
	for i := 0; i < 10; i++ {
		sum += f()
	}
	fmt.Println(sum)
}

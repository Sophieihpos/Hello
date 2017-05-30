package main

import "fmt"

func main() {
	for i := 33; i < 127; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}
}

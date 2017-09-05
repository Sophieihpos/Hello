package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var d []string
	d = strings.Fields(s)
	fmt.Printf("New strings: %s,\n", d)
	l := len(d)

	f := make(map[string]int)

	//for i := 0; i <= l; i++ {
	//f[d[i]] = 1
	//}

	return f
}

func main() {
	var m string
	m = "I am learning go"
	fmt.Println(WordCount(m))
	wc.Test(WordCount)
}

package main

func main() {
	x := fibo()
	y := nacci(x)
}

func fibo() int {
	a, b := 0, 1
	return int {
		a, b := b, a+b
		return a
	}
}

func nacci(x int) chan int {
	out := make(chan int)
	go func() {
			out <- x
		}
		close(out)
	}()
	return out
}

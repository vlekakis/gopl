package main

import (
	"flag"
	"fmt"
)

func fibo(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	var n = flag.Int("n", 1, "Value of n to compute fibo")
	flag.Parse()
	res := fibo(*n)

	fmt.Printf("The %dth fibo number is %d\n", *n, res)
}

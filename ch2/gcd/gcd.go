package main

import (
	"flag"
	"fmt"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var x = flag.Int("x", 1, "first number")
	var y = flag.Int("y", 0, "second number")
	flag.Parse()
	res := gcd(*x, *y)
	fmt.Printf("The gcd(%d, %d) is %d\n", *x, *y, res)
}

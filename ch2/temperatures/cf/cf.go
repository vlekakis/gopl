package main

import (
	"fmt"
	"os"
	"strconv"

	"Temperatures/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n",
			f, tempconv.FToC(f), f, tempconv.FToK(f),
			c, tempconv.CToF(c), c, tempconv.CToK(c),
			k, tempconv.KToF(k), k, tempconv.KToC(k))

	}
}

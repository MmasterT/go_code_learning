package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	excercise_2()
}

//Excercise 1 was creating the package tempconv (in this directory and $GOPAATH in my desktop)

//todo excercise 2. Change Celisus and Kelvin types to feet and meters. Initialize the const values as in ../packages/tempconv.go.
//Then make the progrram to read command line arguments and os.Stdin.

func excercise_2() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		k := tempconv.Kelvin(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			k, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

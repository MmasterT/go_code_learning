package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	excercise_2()
}

//Excercise 1 was creating the package tempconv (in this directory and $GOPAATH in my desktop)

// excercise 2. Change Celisus and Kelvin types to feet and meters. Initialize the const values as in ../packages/tempconv.go.
//Then make the progrram to read command line arguments and os.Stdin.

type Meter float32
type Feet float32

func distM2F(m Meter) Feet {
	return Feet(m * 3.28084)
}

func distF2M(f Feet) Meter {
	return Meter(f / 3.28084)
}

func excercise_2() {

	if len(os.Args[1:]) != 0 {

		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "mf: %v\n", err)
				os.Exit(1)
			}
			m := Meter(t)
			f := Feet(t)
			fmt.Printf("%v = %v, %v = %v\n",
				m, distM2F(m), f, distF2M(f))
		}

	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			temp, _ := strconv.ParseFloat(input.Text(), 32)
			m := Meter(temp)
			f := Feet(temp)
			fmt.Printf("%v = %v meters, %v = %v foots\n",
				m, distM2F(m), f, distF2M(f))
		}
	}
}

//exercsie 3 done
var sum byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func excercise_3(){
	
	func PopCount( x uint64) int {
		
		
		for i := range 8 {
			sum = sum + pc[byte(x>>i*8)]
		}
		return int(sum)
	}
}

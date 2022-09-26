package main

//echo3 prints the arguments in one line separated by a blank space. This version of echo does not collect the garbage of iterating over a string a redefining a string.

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

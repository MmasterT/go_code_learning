// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //here I defined both variables and they are initialized as '0 value vavriables' 0 for ints, "" for strings
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

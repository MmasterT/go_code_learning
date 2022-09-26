// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// in this function we call the function countLines() before its declaration. Functions and other package level entities may be declared in any order.

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// notice how in the for loop there iss two var, f is the file and err is the error. error == nil when the file is opend correctly, otherwise it states the error. Here we do not manage the error but we print it.
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			//when a funcction recives a "map" it recives a copy of the reference, so any changes  the function makes to the underlaying data will be visible to the map callers.
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

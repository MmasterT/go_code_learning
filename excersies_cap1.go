package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	excercise_1()
	excercise_2()
}

func excercise_1() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

func excercise_2() {
	start := time.Now()
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + arg + " " + strconv.Itoa(index)
		sep = "\n"
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

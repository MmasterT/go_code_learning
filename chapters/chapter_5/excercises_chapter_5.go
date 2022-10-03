package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//excercise_15()
	excercise_16()
}

func excercise_15() {
	if len(os.Args) < 2 {
		log.Fatal("You have to use at least one argument as an input.")
	}

	args := os.Args[1:]
	var vals []int

	for i := 0; i < len(args); i++ {
		value := args[i]
		v, _ := strconv.Atoi(value)
		vals = append(vals, v)
	}
	fmt.Printf("the minumum value is :%d\n", min(vals...))
	fmt.Printf("the maximum value is :%d\n", max(vals...))
}

func min(values ...int) int {
	min_value := 0
	for i := 1; i < len(values); i++ {
		min_value := values[0]
		if min_value > values[i] {
			min_value = values[i]
		}
	}

	return min_value
}

func max(values ...int) int {
	max_value := 0
	for i := 1; i < len(values); i++ {
		if values[i] > max_value {
			max_value = values[i]
		}
	}
	return max_value
}

func excercise_16() {
	if len(os.Args) < 2 {
		log.Fatal("You have to use at least one argument as an input.")
	}
	args := os.Args[1:]
	fmt.Println(Join(args...))
}

func Join(s ...string) string {
	var strings string
	for i := 0; i < len(s); i++ {
		strings = strings + s[i]
	}
	return strings
}

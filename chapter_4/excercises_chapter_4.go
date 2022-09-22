package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func main() {
	//excercise_1()
	excercise_2()
}

func excercise_1() {
	s1 := "unodostresquatro"
	s2 := "UNODOSTRESQUATRO"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Println(BitsDifference(&h1, &h2))
}

func BitsDifference(h1, h2 *[sha256.Size]byte) int {
	n := 0
	for i := range h1 {
		for b := h1[i] ^ h2[i]; b != 0; b &= b - 1 {
			n++
		}
	}
	return n
}

func excercise_2() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Fprintf(os.Stderr, "hash error: %v\n", err)
	}

	text = strings.TrimSpace(text)

	if len(os.Args[0:]) != 0 {
		arg := os.Args[1]

		if len(os.Args[0:]) > 2 {
			fmt.Print("too many arguments\n")
		}

		switch {

		case arg == "sha384":
			h1 := sha512.Sum384([]byte(text))
			fmt.Printf("%x\n", h1)

		case arg == "sha512":
			h1 := sha512.Sum512([]byte(text))
			fmt.Printf("%x\n", h1)
		}
	}

	h1 := sha256.Sum256([]byte(text))
	fmt.Printf("%x\n", h1)
}

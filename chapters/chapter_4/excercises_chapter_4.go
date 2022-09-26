package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func main() {
	//excercise_1()
	//excercise_2()
	//excercise_3()
	//excercise_4()
	//excercise_5()
	//TODO excercises 6 and 7
	//excercise_8()
	//TODO excercise 9
	excercise_10()
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

func excercise_3() {
	a := []int{0, 1, 2, 3, 4, 5}
	p := &a
	reverse_pointer(p)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}
func reverse_pointer(s *[]int) {
	for i, j := 0, len((*s))-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func excercise_4() {
	s := []int{0, 1, 2, 3, 4, 5}
	p := &s
	fmt.Println(rotate(p, 2))
}

func rotate(s *[]int, index int) []int {
	for i := range *s {

		if i < index {
			(*s) = append((*s), (*s)[i])
		}
	}

	(*s) = (*s)[index:]
	return *s
}

func excercise_5() {
	s := []string{"5", "6", "7", "7", "8", "9"}
	var p *[]string = &s
	fmt.Println(dedup_adjacen(p))
}

func dedup_adjacen(strings *[]string) []string {
	for i := 1; i < len(*strings); i++ {
		if (*strings)[i-1] == (*strings)[i] {
			copy((*strings)[i:], (*strings)[i+1:])
			(*strings) = (*strings)[:len(*strings)-1]
		}
	}
	return *strings
}

func excercise_8() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	utfletter := make(map[rune]int)
	utfdigit := make(map[rune]int)
	invalid := 0 // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			utfletter[r]++
		}

		if unicode.IsNumber(r) {
			utfdigit[r]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range utfletter {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("digit\tcount\n")
	for c, n := range utfdigit {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func excercise_10() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n\n", result.TotalCount)

	for _, item := range result.Items {
		time_creation := item.CreatedAt
		t1 := time_creation

		t2 := time.Now()
		if t1.Year() < t2.Year() {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login, item.Title, "created more than a year")
		}
		if t1.Month() < t2.Month() && (t1.Day()-t2.Day()) < 0 {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login, item.Title, "created with more than a year")
		} else {
			fmt.Println("issues with more than a month and less than a year")
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login, item.Title, "created more than a month but less than a year")
		}
	}
}

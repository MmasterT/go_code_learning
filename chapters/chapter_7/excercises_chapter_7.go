package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//excersice_1()
	excercise_3()
}

func excersice_1() {
	input := os.Stdin
	var buf bytes.Buffer
	io.Copy(&buf, input)
	asString := string(buf.Bytes())
	var a WordCounter
	var b LineCounter

	a.Write(asString)
	b.Write(asString)
	fmt.Printf("Word counter: %v\nLine counter: %v\n", a, b)
}

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(input string) (int, error) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	count := 0
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*c += WordCounter(count)
	return count, nil
}

func (y *LineCounter) Write(input string) (int, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)
	count := 0

	for s.Scan() {
		fmt.Println(s.Text())
		count++
	}

	err := s.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*y += LineCounter(count)
	return count, nil
}

func excersice_2() {
	/*
		Exercis e 7.2: Wr ite a function CountingWriter with the sig nature below that, given an
		io.Writer, retur ns a new Writer that wraps the original, and a point er to an int64 var iable
		that at any mom ent contains the number of bytes writt en to the new Writer.
		func CountingWriter(w io.Writer) (io.Writer, *int64)
	*/

}

type byteCounter struct {
	w       io.Writer
	written int64
}

func (c *byteCounter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.written += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.written
}

func excercise_3() {
	/*
	   Exercis e 7.3: Wr ite a String method for the *tree type in gopl.io/ch4/treesort (§4.4)
	   that reveals the sequence of values in the tre e.
	*/
	root := &tree{value: 3}
	root = add(root, 2)
	root = add(root, 4)
	fmt.Println(root.String())
}

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	order := make([]int, 0)
	order = appendValues(order, t)
	if len(order) == 0 {
		return "[]"
	}

	b := &bytes.Buffer{}
	b.WriteByte('[')
	b.WriteByte(' ')
	for _, v := range order[0:] {
		fmt.Fprintf(b, "%d ", v)
	}
	b.WriteByte(']')
	return b.String()
}

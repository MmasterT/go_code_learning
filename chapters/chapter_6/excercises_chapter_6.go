package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popcount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func main() {
	excersice_1()
}

func excersice_1() {
	/*
	   Exercise 6.1: Implement these addition al methods:
	   func (*IntSet) Len() int // return the number of elements
	   func (*IntSet) Remove(x int) // remove x from the set
	   func (*IntSet) Clear() // remove all elements from the set
	   func (*IntSet) Copy() *IntSet // return a copy of the set
	*/
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Printf("Length after removing number 9\nx = %v\ny = %v\n", x.Len(), y.Len())
	x.Remove(9)
	y.Remove(9)
	fmt.Println("removing the number 9...")
	fmt.Println(y.String())
	fmt.Println(x.String())
	fmt.Printf("Length after removing number 9\nx = %v\ny = %v\n", x.Len(), y.Len())
	fmt.Printf("Length after removing number 9\nx = %v\ny = %v\n", x.Len(), y.Len())

	z := (&x).Copy()
	fmt.Printf("Copy of y\ny = z = %v\n", z)

}

// return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	new := &IntSet{}
	new.words = make([]uint64, len(s.words))
	copy(new.words, s.words)
	return new
}

func excersice_2() {
	/*
		Exercise 6.2: Define a variadic (*IntSet).AddAll(...int) method that allows a list of values
		to be added, such as s.AddAll(1, 2, 3).
	*/
}

func (x *IntSet) AddAll(ints ...int) {
	for _, i := range ints {
		x.Add(i)
	}
}

func excercise_3() {
	/*
		(*IntSet).UnionWith comp utes the union of two sets using |, the word-p arallel
		bit w ise OR operator. Imp lement met hods for IntersectWith, DifferenceWith, and SymmetricDifference
		for the cor responding set operat ions. (The symmetr ic dif ference of two sets cont ains the elements present in one set or the other but not both.)
	*/

}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

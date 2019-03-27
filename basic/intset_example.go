package main

import (
	"fmt"
)

type Intset struct {
	words []int64
}

func (s *Intset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return len(s.words) > word && (s.words[word]&(1<<bit) != 0)
}

func (s *Intset) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

func (s *Intset) UnionWith(t *Intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	var s, t Intset
	s.Add(3)
	s.Add(21)
	t.Add(5)
	t.Add(86)
	t.Add(12)

	fmt.Println("s =", s)
	fmt.Println("t =", t)

	if s.Has(3) {
		fmt.Println("s has", 3)
	}

	if s.Has(5) {
		fmt.Println("s has not", 5)
	}

	s.UnionWith(&t)
	fmt.Println("s union t =", s)
}

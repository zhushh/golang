package main

import (
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, v int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = v
		return t
	}
	if t.value < v {
		t.right = add(t.right, v)
	} else {
		t.left = add(t.left, v)
	}
	return t
}

func main() {
	var val = []int{3, 12, 3, 5, 18, 9}
	fmt.Println(val)
	Sort(val)
	fmt.Println(val)
}

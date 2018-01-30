// https://tour.golang.org/concurrency/8
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

/* my  go
func Walk(t *tree.Tree, ch chan int) {
	if reflect.TypeOf(t.Left).String() == "Tree" {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if reflect.TypeOf(t.Right).String() == "Tree" {
		Walk(t.Right, ch)
	}
}
*/

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

/* my go
func Same(t1, t2 *tree.Tree) bool {
	var s []int

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	close(ch1)
	close(ch2)

	for _, val := range ch1 {
		append(s, val)
	}

	for _, v := range ch2 {
		var included bool

		for i, w := range s {
			if w == v {
				s[i] = 0
				included = true
			}
		}

		if !included {
			return false
		}
	}

	return true
}
*/

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	// close(ch)

	for val := range ch {
		fmt.Println(val)
	}

	result1 := Same(tree.New(1), tree.New(1))
	result2 := Same(tree.New(1), tree.New(2))

	fmt.Printf("Result 1: %v, Result 2: %v", result1, result2)
}

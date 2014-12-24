package main

import (
	"fmt"
	"github.com/openminder/go-tutorials/stacker/stack"
)

func main() {
	var haystack stack.Stack
	// push different types of objects into the stack. the stack interface allows any type
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)
	// Print items till stack is empty - for loop without params runs infinit till break
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World!"
	// The first argument is the command name, so check if the Args is bigger then 1
	if len(os.Args) > 1 {
		// Join all strings after the command. e.g. "hello John, Jim and Frank"
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}

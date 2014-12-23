package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ",
		"11 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		"111"},
	{" 222 ",
		"2   2 ",
		"2   2",
		"   2 ",
		" 2   ",
		"2    ",
		"22222"},
}

func main() {
	// os.Args has a minimum length of one because the go command counts
	// So to check if there any args we can check if length is 1
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s: <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// Store first argument that the user enters as a string
	stringOfDigits := os.Args[1]
	// go throw each line and print it on the screen if it's in the correct (0-2) range
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			// When indexing a position of a string we get the utf8 value of it
			// convert the utf value into the byte value
			// e.g. 3 = utf-8(codepoint51) 0 = utf-8(codepoint 48) so 51 - 48 = 3
			digit := stringOfDigits[column] - '0'
			// check if digit is in the provided range
			if 0 <= digit && digit <= 2 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}

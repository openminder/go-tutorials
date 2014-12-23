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
	}

	// Store first argument that the user enters as a string
	stringOfDigits := os.Args[1]
	// go throw each line and print it on the screen if it's in the correct (0-2) range
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 2 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}

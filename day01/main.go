package main

import (
	"fmt"

	"github.com/srowles/adventofcode2015"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2015.MustInputFromWebsite("1")
	fmt.Println(floor(input))
}

func do2() {
	input := adventofcode2015.MustInputFromWebsite("1")
	fmt.Println(minus1Floor(input))
}

func floor(input string) int {
	floor := 0
	for _, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func minus1Floor(input string) int {
	floor := 0
	charPos := 1
	for _, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return charPos
		}
		charPos++
	}
	return -1
}

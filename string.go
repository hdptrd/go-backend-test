package main

import (
	"fmt"
)

// declare custom variable type
type inputString string

func (inp inputString) print() {
	fmt.Print(inp)
}

func (inp inputString) findIndexOfChar(char string) int {
	var index int
	for i, r := range inp {
		if string(r) == char {
			index = i
			break
		}
	}
	return index
}

func (inp inputString) getStringStartFromIndex(idx int) string {
	var str string
	for i, r := range inp {
		if i >= idx {
			str += string(r)
		}
	}
	return str
}

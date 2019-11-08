package main

import (
	"fmt"
	"os"
)

var usPhone = map[byte]int{
	'a': 2,
	'b': 2,
	'c': 2,
	'd': 3,
	'e': 3,
	'f': 3,
	'g': 4,
	'h': 4,
	'i': 4,
	'j': 5,
	'k': 5,
	'l': 5,
	'm': 6,
	'n': 6,
	'o': 6,
	'p': 7,
	'q': 7,
	'r': 7,
	's': 7,
	't': 8,
	'u': 8,
	'v': 8,
	'w': 9,
	'x': 9,
	'y': 9,
	'z': 9,
	' ': 0,
}

func dial(s string, m map[byte]int) (result []int) {
	for _, b := range []byte(s) {
		if i, ok := m[b]; ok {
			result = append(result, i)
		}
	}
	return
}

func main() {
	for _, arg := range os.Args[1:] {
		ns := dial(arg, usPhone)
		if len(ns) > 0 {
			fmt.Printf("Dialing %q: %v\n", arg, ns)
		} else {
			fmt.Printf("Nothing to dial for %q\n", arg)
		}
	}
}

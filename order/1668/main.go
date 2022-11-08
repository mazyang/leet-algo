package main

import "strings"

func maxRepeating(sequence string, word string) int {
	res := 1
	tmp := word
	for strings.Contains(sequence, tmp) {
		res++
		tmp = tmp + word
	}
	return res - 1
}

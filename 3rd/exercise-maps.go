package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		count, ok := m[word]
		if ok {
			m[word] = count + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func exercise_maps_main() {
	wc.Test(WordCount)
}

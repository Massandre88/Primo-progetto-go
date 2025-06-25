package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mappa := make(map[string]int)
	var stringa string
	for _, c := range s {
		if c == ' ' {
			mappa[stringa]++
			stringa = ""
		} else {
			stringa += string(c)
		}
	}
	mappa[stringa]++
	return mappa
}

func main() {
	wc.Test(WordCount)
}

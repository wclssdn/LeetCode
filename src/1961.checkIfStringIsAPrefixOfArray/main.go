package main

import "fmt"

func main() {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcode", "apples"}
	r := isPrefixString(s, words)
	fmt.Println("result", r)
}

func isPrefixString(s string, words []string) bool {
	start := 0
	total := len(s)
	for _, word := range words {
		l := len(word)
		if start+l > total {
			return false
		}
		if s[start:start+l] != word {
			return false
		}
		start += l
		if start == total {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
)

func main() {
	s := "()[{}]"
	r := isValid(s)
	fmt.Println(s, r)
}

func isValid(s string) bool {
	slice := make([]uint8, 0, 100)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			slice = append(slice, s[i])
		} else {
			if len(slice) == 0 {
				return false
			}
			switch s[i] {
			case ']':
				if slice[len(slice)-1] != '[' {
					return false
				}
			case '}':
				if slice[len(slice)-1] != '{' {
					return false
				}
			case ')':
				if slice[len(slice)-1] != '(' {
					return false
				}
			}
			slice = slice[:len(slice)-1]
		}
	}

	return len(slice) == 0
}

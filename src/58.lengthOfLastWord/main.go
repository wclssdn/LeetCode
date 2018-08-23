package main

import "fmt"

func main() {
	str := "hello world "
	r := lengthOfLastWord(str)
	fmt.Println("lenth of last world is", r)
}
func lengthOfLastWord(s string) int {
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if r == 0 {
				continue
			}
			break
		}
		r++
	}
	return r
}

package main

import "fmt"

func main() {
	s := []string{"flower", "flow", "flight", ""}
	r := longestCommonPrefix(s)
	fmt.Println(s, r)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	p := ""
Loop:
	for i := 0; ; i++ {
		char := 0
		for x := range strs {
			if len(strs[x]) < i+1 {
				break Loop
			}
			if char == 0 {
				char = int(strs[x][i])
			} else if char != int(strs[x][i]) {
				break Loop
			}
		}
		p += string(char)
	}
	return p
}

package main

import "fmt"

func main() {
	roman := "LVIII"

	r := romanToInt(roman)

	fmt.Println(roman, r)
}

func romanToInt(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	r := 0
	last := 0
	for i := 0; i < len(s); i++ {
		tmp := m[string(s[i])]
		if last > 0 && last < tmp {
			r += tmp - last*2
			last = 0
		} else {
			r += tmp
			last = tmp
		}
	}
	return r
}

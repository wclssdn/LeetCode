package main

import (
	"fmt"
	"math"
)

func main() {
	a := "1111"
	b := "1111"

	r := addBinary(a, b)
	fmt.Println(a, "+", b, "=", r)
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	max := int(math.Max(float64(la), float64(lb)))
	r := make([]rune, max+1)
	carry := false
	var _a uint8 = 0
	var _b uint8 = 0
	for i := 0; i < max; i++ {
		_a = 0
		if i < la {
			_a = a[la-i-1]
		}
		_b = 0
		if i < lb {
			_b = b[lb-i-1]
		}
		if _a == '0' && _b == '0' {
			if carry {
				r[max-i] = '1'
				carry = false
			} else {
				r[max-i] = '0'

			}
		} else if _a == '0' || _b == '0' {
			if carry {
				if _a == 0 || _b == 0 {
					r[max-i] = '1'
					carry = false
				} else {
					r[max-i] = '0'
				}
			} else {
				if _a == 0 || _b == 0 {
					r[max-i] = '0'
				} else {
					r[max-i] = '1'
				}
			}
		} else {
			if carry {
				if _a == 0 || _b == 0 {
					r[max-i] = '0'
				} else {
					r[max-i] = '1'
				}
			} else {
				if _a == 0 || _b == 0 {
					r[max-i] = '1'
				} else {
					r[max-i] = '0'
					carry = true
				}
			}
		}
	}
	if carry {
		r[0] = '1'
		return string(r)
	}
	return string(r[1:])
}

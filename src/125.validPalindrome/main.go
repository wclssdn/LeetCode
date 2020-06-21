package main

import "fmt"

/**
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/
func main() {
	s := "A man, a plan, a canal: Panama"
	//s := "race a car"
	//s := "0P"
	isPalindrome := isPalindrome(s)
	fmt.Println("is isPalindrome:", isPalindrome)
}
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	l := len(s) - 1
	if l == 0 {
		return true
	}

	for i := 0; i <= l; i++ {
		if i == l {
			print(1)
			return true
		}
		if !isLetter(s[i]) {
			continue
		}
		if !isLetter(s[l]) {
			l--
			i--
			continue
		}
		if !isSame(s[i], s[l]) {
			return false
		}
		l--
	}
	return true
}

func isLetter(b byte) bool {
	return isAlpha(b) || isNum(b)
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSame(a, b byte) bool {
	if isNum(a) || isNum(b) {
		return a == b
	}
	return a == b || a == b-'A'+'a' || a == b+'A'-'a'
}

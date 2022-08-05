package main

import "fmt"

func main() {
	s := "aba"
	fmt.Println(palindrome(s))
}

func palindrome(s string) bool {
	l := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[l-i] {
			return false
		}
	}
	return true
}

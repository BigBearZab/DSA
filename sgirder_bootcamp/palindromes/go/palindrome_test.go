package main

import "testing"

func TestPalindrome(t *testing.T) {
	inputs := []string{"aba", "cba", "ogra", "aaaaaaaaaaaaaaa"}
	outputs := []bool{true, false, false, true}

	for i, _ := range inputs {
		if palindrome(inputs[i]) != outputs[i] {
			t.Errorf(
				"Expected %v, for palindrome of %v, but got %v", outputs[i], inputs[i], palindrome(inputs[i]),
			)
		}
	}
}

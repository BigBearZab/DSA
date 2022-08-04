package main

import (
	"testing"
)

func TestStrRev(t *testing.T) {
	test_str := "abc"
	test_out := strRev(test_str)
	if test_out != "cba" {
		t.Errorf("Expected %v but got %v instead", "cba", test_out)
	}
}

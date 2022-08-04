package main

import "fmt"

func main() {
	fmt.Println(strRev("some str"))
}

func strRev(s string) string {
	o := ""
	for i := 0; i < len(s); i++ {
		o = string(s[i]) + o
	}
	return o
}

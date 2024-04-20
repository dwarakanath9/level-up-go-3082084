package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(" this is the test msg")

	msg := "this is text string"

	words := strings.Split(msg, " ")
	fmt.Println("words :", words)
	fmt.Println("msg :", msg)
	rp := strings.Repeat("msg", 10)
	fmt.Println("stirng is repeating :", rp)

	for _, w := range words {
		var s []string
		for i, c := range w {
			sr := strings.Repeat(string(c), i+1)
			s = append(s, sr)
		}
		fmt.Println("s[]", s)
		joinstring := strings.Join(s, "")
		fmt.Println("joinstring:", joinstring)
	}

}

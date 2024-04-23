package main

import "fmt"

func main() {
	s := "hello world"
	cm := make(map[rune]int)

	for _, st := range s {
		cm[st]++
	}

	maxChar := ' '
	maxCount := 0

	for c, count := range cm {
		if count > maxCount {
			maxChar = rune(c)
			maxCount = count

		}
	}

	fmt.Printf("max char %c and max count %d ", maxChar, maxCount)
}

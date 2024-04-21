package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.20
	v := 0.99
	fmt.Println("Floor of", x, "is", math.Floor(v/x)) // Output: Floor of 3.6 is 3
	fmt.Println("without Floor of", x, "is", (v / x)) // Output: Floor of 3.6 is 3

}

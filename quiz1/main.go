package main

import (
	"fmt"
)

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, tempInt := range intSlice {
		if tempInt%2 == 0 {
			fmt.Printf("%v is even\n", tempInt)
		} else {
			fmt.Printf("%v is odd\n", tempInt)
		}

	}
}

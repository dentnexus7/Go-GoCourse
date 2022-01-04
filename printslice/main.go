package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, _ := range numbers {
		if numbers[i]%2 == 0 {
			fmt.Printf("%d is even\n", numbers[i])
		} else {
			fmt.Printf("%d is odd\n", numbers[i])
		}
	}
}

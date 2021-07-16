package main

import "fmt"

func main() {
	limit := 0
	fmt.Scanf("%d", &limit)

	for iter := 0; iter < limit; iter++ {
		if iter == (limit - 1) {
			fmt.Printf("Ho!\n")
		} else {
			fmt.Printf("Ho ")
		}
	}
}

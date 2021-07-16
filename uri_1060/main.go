package main

import "fmt"

func main() {
	positiveCount, number := 0, 0.0

	for iter := 0; iter < 6; iter++ {
		fmt.Scanf("%f", &number)

		if  number >= 0 {
			positiveCount += 1
		}
	}

	fmt.Printf("%d valores positivos\n", positiveCount)
}
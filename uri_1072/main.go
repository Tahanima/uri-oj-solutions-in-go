package main

import (
	"fmt"
)

func main()  {
	limit, number, countOfNumbersIn, countOfNumbersOut := 0, 0, 0, 0
	fmt.Scanf("%d", &limit)

	for iter := 0; iter < limit; iter++ {
		fmt.Scanf("%d", &number)

		if number >= 10 && number <= 20 {
			countOfNumbersIn++
		} else {
			countOfNumbersOut++
		}
	}

	fmt.Printf("%d in\n%d out\n", countOfNumbersIn, countOfNumbersOut)
}

package main

import "fmt"

func main()  {
	number := 0
	fmt.Scanf("%d", &number)

	if number % 2 == 0 {
		fmt.Printf("%d\n", number + 2)
	} else {
		fmt.Printf("%d\n", number + 1)
	}
}

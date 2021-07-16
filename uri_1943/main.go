package main

import "fmt"

func main()  {
	position := 0
	fmt.Scanf("%d", &position)

	if position == 1 {
		fmt.Printf("Top 1\n")
	} else if position > 1 && position <= 3 {
		fmt.Printf("Top 3\n")
	} else if position > 3 && position <= 5 {
		fmt.Printf("Top 5\n")
	} else if position > 5 && position <= 10 {
		fmt.Printf("Top 10\n")
	} else if position > 10 && position <= 25 {
		fmt.Printf("Top 25\n")
	} else if position > 25 && position <= 50 {
		fmt.Printf("Top 50\n")
	} else {
		fmt.Printf("Top 100\n")
	}
}

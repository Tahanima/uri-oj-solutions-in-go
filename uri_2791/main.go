package main

import "fmt"

func main()  {
	countOfBean := 0

	for iter := 1; iter <= 4; iter++ {
		fmt.Scanf("%d", &countOfBean)

		if countOfBean == 1 {
			fmt.Printf("%d\n", iter)
			break
		}
	}
}

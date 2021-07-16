package main

import "fmt"

func main()  {
	evenCount, oddCount, positiveCount, negativeCount, number := 0, 0, 0, 0, 0

	for iter := 0; iter < 5; iter++ {
		fmt.Scanf("%d", &number)

		if number > 0 {
			positiveCount++
		} else if number < 0 {
			negativeCount++
			number *= -1
		}

		if number % 2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Printf("%d valor(es) par(es)\n" +
		"%d valor(es) impar(es)\n" +
		"%d valor(es) positivo(s)\n" +
		"%d valor(es) negativo(s)\n",
		evenCount, oddCount, positiveCount, negativeCount)
}

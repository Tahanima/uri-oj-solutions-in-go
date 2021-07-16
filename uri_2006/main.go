package main

import "fmt"

func main()  {
	teaType, answerOfContestant, countOfCorrectAnswers := 0, 0, 0

	fmt.Scanf("%d", &teaType)

	for iter := 0; iter < 5; iter++ {
		fmt.Scanf("%d", &answerOfContestant)
		if teaType == answerOfContestant {
			countOfCorrectAnswers++
		}
	}

	fmt.Printf("%d\n", countOfCorrectAnswers)
}

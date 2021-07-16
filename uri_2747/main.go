package main

import (
	"fmt"
)

func main()  {
	line0, line1 := "---------------------------------------", "|                                     |"

	for iter := 0; iter < 7; iter++ {
		if iter == 0 || iter == 6 {
			fmt.Println(line0)
		} else {
			fmt.Println(line1)
		}
	}
}

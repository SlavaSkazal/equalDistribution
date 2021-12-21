package main

import (
	"fmt"
)

func main() {
	//подразумевается, что входные данные корректные, проверку не провожу
	var cups [10]int

	var totalLiters, shareOfCup, numberOfTransfusions int

	for i, _ := range cups {
		fmt.Scanf("%d\n", &cups[i])
		totalLiters += cups[i]
	}

	shareOfCup = totalLiters / 10

	for _, v := range cups {
		if v > shareOfCup {
			numberOfTransfusions++
		}
	}

	fmt.Println(numberOfTransfusions)
}

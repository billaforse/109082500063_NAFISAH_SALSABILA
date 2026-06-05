package main

import "fmt"

func main() {
	var x int
	var total, valid int
	var suara [21]int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			suara[x]++
			valid++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", valid)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}

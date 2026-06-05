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

	ketua := 1
	wakil := 1

	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && suara[i] > suara[wakil] {
			wakil = i
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}

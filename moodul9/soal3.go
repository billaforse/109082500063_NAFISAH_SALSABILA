package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A, &B)

	var a, b, i int
	i = 1

	for {
		fmt.Scan(&a, &b)
		if a < 0 || b < 0 {
			break
		}

		fmt.Print("Hasil ", i, " : ")
		if a > b {
			fmt.Println(A)
		} else if b > a {
			fmt.Println(B)
		} else {
			fmt.Println("Draw")
		}
		i++
	}
	fmt.Println("Pertandingan selesai")
}

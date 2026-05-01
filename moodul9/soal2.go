package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idx, cari int
	fmt.Scan(&n)

	a := make([]int, n)
	sum := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()

	fmt.Scan(&idx)
	for i := 0; i < n; i++ {
		if i != idx {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()

	avg := float64(sum) / float64(n)
	fmt.Println(avg)

	var v float64
	for _, val := range a {
		v += math.Pow(float64(val)-avg, 2)
	}
	fmt.Println(math.Sqrt(v / float64(n)))

	fmt.Scan(&cari)
	f := 0
	for _, val := range a {
		if val == cari {
			f++
		}
	}
	fmt.Println(f)
}

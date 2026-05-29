package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	var total float64 = 0

	for i := 0; i < n; i++ {
		total += arr[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max, avg float64

	fmt.Print("Masukkan banyak data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Berat ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	avg = rerata(data, n)

	fmt.Println("Minimum:", min)
	fmt.Println("Maksimum:", max)
	fmt.Println("Rata-rata:", avg)
}

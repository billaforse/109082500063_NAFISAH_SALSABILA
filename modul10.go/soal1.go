package main

import "fmt"

type arrKelinci [1000]float64

func minMax(arr arrKelinci, n int, min *float64, max *float64) {
	*min = arr[0]
	*max = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *min {
			*min = arr[i]
		}
		if arr[i] > *max {
			*max = arr[i]
		}
	}
}

func main() {
	var data arrKelinci
	var n int
	var min, max float64

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Berat kelinci ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	minMax(data, n, &min, &max)

	fmt.Println("Berat minimum:", min)
	fmt.Println("Berat maksimum:", max)
}

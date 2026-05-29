package main

import "fmt"

type arrIkan [1000]float64

func main() {
	var data arrIkan
	var x, y int
	var total float64 = 0
	var idx int = 0

	fmt.Print("Masukkan jumlah wadah dan ikan per wadah: ")
	fmt.Scan(&x, &y)

	n := x * y

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan berat ikan ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	fmt.Println("Total berat per wadah:")

	for i := 0; i < x; i++ {
		var sum float64 = 0
		for j := 0; j < y; j++ {
			sum += data[idx]
			idx++
		}
		fmt.Println("Wadah", i+1, ":", sum)
		total += sum
	}

	rata := total / float64(x)
	fmt.Println("Rata-rata per wadah:", rata)
}

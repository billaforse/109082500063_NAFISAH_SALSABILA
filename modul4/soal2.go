package main

import "fmt"

func hitungSkor(soal, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama string
	var soal, skor int

	var maxNama string
	var maxSoal, minSkor int
	first := true

	for {
		fmt.Scan(&nama)
		if nama == "Selesai." {
			break
		}

		hitungSkor(&soal, &skor)

		if first || soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxNama = nama
			maxSoal = soal
			minSkor = skor
			first = false
		}
	}

	fmt.Println(maxNama, maxSoal, minSkor)
}

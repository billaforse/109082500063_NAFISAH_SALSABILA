# <h1 align="center">Laporan Praktikum Modul 1 - Prosedur </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### soal1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul10/output/output-soal1.png)
Program ini mencari min & max berat kelinci

### 2.Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.
#### soall2.go

```go
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


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul10/output/output-soal2.png)
Program ini Menghitung total & rata-rata ikan per wadah

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul10/output/output-soal3.png)
Program digunakan untuk mencari min, max, dan rata-rata balita

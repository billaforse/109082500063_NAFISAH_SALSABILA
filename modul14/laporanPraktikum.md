# <h1 align="center">Laporan Praktikum Modul 14 - sorting </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.
#### soal1.go

```go
package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		selectionSort(arr)

		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul14/output/output-soal1.png)
Program ini menyusun nomor-nomor rumah kerabat Hercules secara terurut membesar menggunakan algoritma selection sort..
### 2.Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.
#### soal2.go

```go
package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		var ganjil []int
		var genap []int

		for _, v := range arr {
			if v%2 == 1 {
				ganjil = append(ganjil, v)
			} else {
				genap = append(genap, v)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		for _, v := range ganjil {
			fmt.Print(v, " ")
		}
		for _, v := range genap {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul14/output/output-soal2.png)
Program ini menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.

### 3. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.
#### soal3.go

```go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var x int
	var arr []int

	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		arr = append(arr, x)
	}

	insertionSort(arr)

	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	if len(arr) < 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	selisih := arr[1] - arr[0]
	tetap := true

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != selisih {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul14/output/output-soal4.png)
Program yang digunakan untuk membaca data integer.

### 4. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
#### soal4.go

```go
package main

import "fmt"

type Buku struct {
	id, eksemplar, tahun, rating int
	judul, penulis, penerbit     string
}

// INSERTION SORT (descending by rating)
func insertionSort(arr []Buku) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j].rating < key.rating {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// CETAK 1 FAVORIT
func cetakFavorit(arr []Buku) {
	fmt.Println("Buku Favorit:")
	fmt.Println(arr[0].judul, arr[0].penulis, arr[0].penerbit, arr[0].tahun)
}

// CETAK TOP 5
func cetak5Teratas(arr []Buku) {
	fmt.Println("5 Buku Teratas:")
	limit := 5
	if len(arr) < 5 {
		limit = len(arr)
	}
	for i := 0; i < limit; i++ {
		fmt.Println(arr[i].judul)
	}
}

// BINARY SEARCH (karena sudah di-sort)
func cariBuku(arr []Buku, target int) {
	left, right := 0, len(arr)-1
	found := false

	for left <= right {
		mid := (left + right) / 2

		if arr[mid].rating == target {
			fmt.Println("Ditemukan:")
			fmt.Println(arr[mid].judul, arr[mid].penulis, arr[mid].penerbit, arr[mid].tahun)
			found = true
			break
		} else if arr[mid].rating < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	buku := make([]Buku, n)

	// INPUT DATA
	for i := 0; i < n; i++ {
		fmt.Scan(
			&buku[i].id,
			&buku[i].judul,
			&buku[i].penulis,
			&buku[i].penerbit,
			&buku[i].eksemplar,
			&buku[i].tahun,
			&buku[i].rating,
		)
	}

	// SORT
	insertionSort(buku)

	// OUTPUT
	cetakFavorit(buku)
	cetak5Teratas(buku)

	// CARI
	var r int
	fmt.Scan(&r)
	cariBuku(buku, r)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul14/output/output-soal4.png)
Program digunakan untuk mengelola data buku di dalam suatu perpustakaan.
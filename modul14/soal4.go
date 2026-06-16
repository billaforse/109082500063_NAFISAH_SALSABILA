//Insertion sort 2

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

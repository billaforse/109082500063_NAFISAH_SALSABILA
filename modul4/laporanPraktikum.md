# <h1 align="center">Laporan Praktikum Modul 1 - Prosedur </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n!(n−r)!, sedangkan C(n, r) = n! r!(n−r)!
#### soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var p, k int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p)
	combination(a, c, &k)
	fmt.Println(p, k)

	permutation(b, d, &p)
	combination(b, d, &k)
	fmt.Println(p, k)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul1/output/output-soal1.png)
Program ini menghitung permutasi dan kombinasi

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### soall2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul1/output/output-soall2.png)
Program ini Menghitung skor lomba (jumlah soal + waktu)

### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah 1⁄2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah: 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal. prosedure cetakDeret(in n : integer )masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000. Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.
#### soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n)
		if n == 1 {
			break
		}
		fmt.Print(" ")

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul1/output/output-soal3.png)
Program mencetak deret sampai 1

# <h1 align="center">Laporan Praktikum Modul 12 - searching </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT. Buatlah program pilkart yang akan membaca, memvalidasi, dan menghitung suara yang diberikan dalam pemilihan ketua RT tersebut.
#### soal1.go

```go
package main

import "fmt"

func main() {
	var x int
	var total, valid int
	var suara [21]int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			suara[x]++
			valid++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", valid)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul12/output/output-soal1.png)
Program ini membaca input angka sampai ketemu 0, lalu meng-input suara masuk (Validasi dan menghitung frekuensi).
### 2.Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan nomor peserta terkecil berikutnya.
#### soall2.go

```go
package main

import "fmt"

func main() {
	var x int
	var total, valid int
	var suara [21]int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			suara[x]++
			valid++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", valid)

	ketua := 1
	wakil := 1

	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && suara[i] > suara[wakil] {
			wakil = i
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul12/output/output-soal2.png)
Program ini membaca input angka sampai ketemu 0, lalu meng-input suara masuk dan juga mencari suara terbanyak ketua lalu diikuti wakil (mencari nilai terbesar dan terbesar kedua).

### 3. Diberikan n data integer positif dalam keadaan terurut membesar dan sebuah integer lain k, apakah bilangan k tersebut ada dalam daftar bilangan yang diberikan? Jika ya, berikan indeksnya, jika tidak sebutkan "TIDAK ADA".
#### soal3.go

```go


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul12/output/output-soal3.png)
Program menggunakan Binary search untuk mencari posisi K (Pencarian cepat di data terurut).
# <h1 align="center">Laporan Praktikum Modul 1 - Prosedur </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(p, c.pusat) <= float64(c.radius)
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}



```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul9/output/output-soal1.png)
Program ini menentukan posisi suatu titik terhadap dua lingkaran.

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
#### soall2.go

```go

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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul9/output/output-soall2.png)
Program ini untuk mengolah data array

#### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
#### soal3.go

```go
package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A, &B)

	var a, b, i int
	i = 1

	for {
		fmt.Scan(&a, &b)
		if a < 0 || b < 0 {
			break
		}

		fmt.Print("Hasil ", i, " : ")
		if a > b {
			fmt.Println(A)
		} else if b > a {
			fmt.Println(B)
		} else {
			fmt.Println("Draw")
		}
		i++
	}
	fmt.Println("Pertandingan selesai")
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul9/output/output-soal3.png)
Program dipakai untuk menentukan hasil pertandingan antara dua klub


### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
#### soal4.go

```go
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)

	for i := n - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
	fmt.Println()

	palin := true
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			palin = false
		}
	}

	if palin {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul9/output/output-soal4.png)
Program memproses teks dan mengecek palindrom

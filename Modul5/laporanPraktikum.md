# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### soal1.go

```go
package main
import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fib(i), " ")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul5/output/output-soal1.png)
Program ini menampilkan deret fibonacci sampai N

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### soall2.go

```go
package main
import "fmt"

func bintang(n int) {
	if n == 0 {
		return
	}
	bintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul5/output/output-soal2.png)
Program ini menampilkan pola segitiga bintang dari 1 sampai N

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main
import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul5/output/output-soal3.png)
Program ini menampilkan semua angka yang bisa membagi N


### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.
#### soal4.go

```go
package main

import "fmt"

func pola(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	pola(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul5/output/output-soal4.png)
Program ini menampilkan angka dari N turun ke 1 lalu balik lagi ke N


### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.
#### soal5.go

```go
package main

import "fmt"

func ganjil(n int) {
	if n == 0 {
		return
	}
	ganjil(n - 1)

	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul5/output/output-soal5.png)
Program ini menampilkan semua bilangan ganjil dari 1 sampai N


### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan. Masukan terdiri dari bilangan bulat x dan y. Keluaran terdiri dari hasil x dipangkatkan y.
#### soal6.go

```go
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(pangkat(x, y))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul5/output/output-soal6.png)
Program ini menghitung hasil x dipangkatkan y



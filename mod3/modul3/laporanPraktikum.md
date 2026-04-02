# <h1 align="center">Laporan Praktikum Modul 1 - Pengenalan Alpro </h1>
<p align="center">NAFISAH SALSABILA - 109082500063</p>

## unguided

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul1/output/output-soal1.png)
Program ini menerima 3 input string, lalu ditampilkan urutan awal. setelah itu, program menukar data menggunakan variabel sementara, sehingga urutannya berubah.

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, 'hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
#### soall2.go

```go
package main

import "fmt"

func main() {

	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {

		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if !(w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul1/output/output-soall2.png)
Program ini menerima empat input warna dari pengguna sebanyak lima kali percobaan. Program kemudian periksa kalau urutan warna sudah sesuai dengan yang di input mulai dari merah, kuning, hijau, dan ungu. kalau sudah sesuai, maka akan ditampilkan BERHASIL: true, kalau belum hasilnya BERHASIL: false.

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal3.go

```go
package main

import "fmt"

func main() {

	var berat int
	var kg, sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/billaforse/109082500063_NAFISAH_SALSABILA/blob/main/modul1/output/output-soal3.png)
Program ini digunakan untung menghitung biaya pengiriman parsel berdasarkan beratnya. Profram terima input berat parsel dalam gram, lalu dihitung jumlah kilogram dan sisa gram nya. setelah semua nya dihitung, program menampilkan detail berat dan total biaya pengiriman.

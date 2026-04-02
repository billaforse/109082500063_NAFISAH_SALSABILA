package main

import "fmt"

func main() {
	fmt.Println("====== menampilkan rangkaian informasi tanpa array ")
	var namaku1, namaku2, namaku3, namaku4, namaku5 string
	namaku1 = "rio"
	namaku2 = "eric"
	namaku3 = "bagas"
	namaku4 = "fairuz"
	namaku5 = "cofa"
	fmt.Println(namaku1)
	fmt.Println(namaku2)
	fmt.Println(namaku3)
	fmt.Println(namaku4)
	fmt.Println(namaku5)

	fmt.Println("====== menampilkan rangkaian informasi dengan array ")
	var namaku [8]string
	namaku[0] = "rio"
	namaku[1] = "eric"
	namaku[2] = "bagas"
	namaku[3] = "fairuz"
	namaku[4] = "cofa"
	namaku[5] = "cofa1"
	namaku[6] = "cofa2"
	namaku[7] = "cofa3"
	for i := 0; i <= 7; i++ {
		fmt.Println(namaku[i])
	}

	fmt.Println("====== menampilkan rangkaian informasi dengan SLICE ")
	aint := []int{3, 4, 5, 5, 7}
	for i := 0; i <= len(aint)-1; i++ {
		fmt.Println(aint[i])
	}
	jumlah_room := len(aint)
	fmt.Println("====== Jumlah room dari slice saat ini adalah ", jumlah_room)
	fmt.Println("====== menampilkan rangkaian informasi dengan SLICE setelah append (ditambahkan value baru)")
	aint = append(aint, 38)
	aint = append(aint, 59)
	aint = append(aint, 28)
	for i := 0; i <= len(aint)-1; i++ {
		fmt.Println(aint[i])
	}
}

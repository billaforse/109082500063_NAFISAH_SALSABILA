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
	var namaku [5]string
	namaku[0] = "rio"
	namaku[1] = "eric"
	namaku[2] = "bagas"
	namaku[3] = "fairuz"
	namaku[4] = "cofa"
	fmt.Println(namaku[1])
	fmt.Println(namaku[4])

}

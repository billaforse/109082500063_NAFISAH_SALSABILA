package main

import "fmt"

func PS(n int) {
	//bentuk demonstrasi dari rekursif

	if n >= 0 {
		fmt.Println(n)
		PS(n - 1)
	}
}

func main() {
	fmt.Println("=== menampilkan loop dengan for-loop")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=== menampilkan loop dengan rekursif")
	PS(10)

}

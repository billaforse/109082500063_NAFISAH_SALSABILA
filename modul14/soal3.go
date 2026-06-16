// Insertion sort
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

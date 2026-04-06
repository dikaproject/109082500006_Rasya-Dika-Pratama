package main

import "fmt"

func cetakGanjil(n int) {
	if n < 1 {
		return
	} else {
		cetakGanjil(n - 1)

		if n%2 != 0 {
			fmt.Print(n, " ")
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	cetakGanjil(n)
	fmt.Println()
}
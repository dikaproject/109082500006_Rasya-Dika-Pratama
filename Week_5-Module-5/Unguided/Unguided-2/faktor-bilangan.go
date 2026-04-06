package main

import "fmt"

func cetakFaktor(n int, i int) {
	if i > n {
		return
	} else {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		cetakFaktor(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	cetakFaktor(n, 1)
	fmt.Println()
}
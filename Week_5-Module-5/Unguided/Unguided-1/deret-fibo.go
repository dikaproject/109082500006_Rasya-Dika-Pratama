package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}
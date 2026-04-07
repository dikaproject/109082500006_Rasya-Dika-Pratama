package main

import (
	"fmt"
)

func CekPassword(password string) string {
	if len(password) < 6 {
		return "Password terlalu pendek"
	} else {
		return "Password valid"
	}
}

func main() {
	var password string
	fmt.Print("Masukkan kata sandi: ")
	fmt.Scan(&password)

	result := CekPassword(password)
	fmt.Println(result)
}

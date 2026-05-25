package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	namaDepan    string
	namaBelakang string
	gol          int
	assist       int
}

func main() {
	var n int
	fmt.Scan(&n)

	var pemain [NMAX]Pemain
	for i := 0; i < n; i++ {
		fmt.Scan(&pemain[i].namaDepan, &pemain[i].namaBelakang, &pemain[i].gol, &pemain[i].assist)
	}

	for i := 1; i < n; i++ {
		temp := pemain[i]
		j := i - 1
		for j >= 0 && (pemain[j].gol < temp.gol || (pemain[j].gol == temp.gol && pemain[j].assist < temp.assist)) {
			pemain[j+1] = pemain[j]
			j--
		}
		pemain[j+1] = temp
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", pemain[i].namaDepan, pemain[i].namaBelakang, pemain[i].gol, pemain[i].assist)
	}
}


package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalWadah [1000]float64
	var idxWadah int = 0
	var beratSementara float64 = 0
	var hitungIkan int = 0

	for i := 0; i < x; i++ {
		beratSementara += ikan[i]
		hitungIkan++


		if hitungIkan == y || i == x-1 {
			totalWadah[idxWadah] = beratSementara
			idxWadah++

			beratSementara = 0
			hitungIkan = 0
		}
	}

	var totalSemua float64 = 0

	for i := 0; i < idxWadah; i++ {
		fmt.Printf("%.2f", totalWadah[i])
		if i < idxWadah-1 {
			fmt.Print(" ")
		}
		totalSemua += totalWadah[i]
	}
	fmt.Println()

	var rataRata float64 = totalSemua / float64(idxWadah)
	fmt.Printf("%.2f\n", rataRata)
}
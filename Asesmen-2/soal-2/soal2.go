package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(arr arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(arr arrayMahasiswa, n int, nim string) int {
	maxNilai := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			if arr[i].nilai > maxNilai {
				maxNilai = arr[i].nilai
			}
		}
	}
	return maxNilai
}

func main() {
	var n int
	var data arrayMahasiswa

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	var nimCari string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(data, n, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(data, n, nimCari)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiTerbesar)
}
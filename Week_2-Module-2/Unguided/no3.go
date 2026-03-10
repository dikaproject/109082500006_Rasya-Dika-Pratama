package main

import "fmt"

func hitungBiayaSisa(sisa, totalKg int) int {
    if totalKg >= 10 {
        return 0 
    }
    if sisa >= 500 {
        return sisa * 5
    }
    return sisa * 15
}

func main() {
    var totalGram int
    fmt.Print("Masukkan total berat (gram): ")
    fmt.Scan(&totalGram)

    kg := totalGram / 1000
    sisa := totalGram % 1000
    
    biayaKg := kg * 10000
    biayaTambahan := hitungBiayaSisa(sisa, kg)
    
    totalBiaya := biayaKg + biayaTambahan

    fmt.Println("===== Detail Perhitungan =====")
    fmt.Printf("Detail berat : %d kg + %d gram\n", kg, sisa)
    fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaTambahan)
    fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
package main

import "fmt"

func cekKabisat(tahun int) bool {
    return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
    var tahun int
    fmt.Print("Masukkan tahun : ")
    fmt.Scan(&tahun)

    status := cekKabisat(tahun)

    fmt.Println("Kabisat:", status)
}
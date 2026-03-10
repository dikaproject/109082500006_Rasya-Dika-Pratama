package main

import "fmt"

func cekUrutanWarna() bool {
    var w1, w2, w3, w4 string
    fmt.Scan(&w1, &w2, &w3, &w4)
    return w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu"
}

func main() {
    berhasil := true

    for i := 1; i <= 5; i++ {
        fmt.Printf("Percobaan %d : ", i)
        if !cekUrutanWarna() {
            berhasil = false
        }
    }

    fmt.Println("BERHASIL :", berhasil)
}